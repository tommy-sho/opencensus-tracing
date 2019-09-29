package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"github.com/kelseyhightower/envconfig"
	"github.com/tommy-sho/opencensus-tracing/user_service/config"
	proto "github.com/tommy-sho/opencensus-tracing/user_service/genproto"
	"github.com/tommy-sho/opencensus-tracing/user_service/service"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	port = "5001"
)

func main() {
	var conf config.Config
	err := envconfig.Process("", &conf)

	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: conf.GetGoogleCloudProject(),
	})
	if err != nil {
		log.Fatal(err)
	}
	trace.RegisterExporter(exporter)

	s := service.NewUserService()
	server := grpc.NewServer()
	proto.RegisterUserServiceServer(server, s)
	RegisterHeathCheck(server)
	reflection.Register(server)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		panic(err)
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan,
		os.Interrupt,
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	go func() {
		<-stopChan
		gracefulStopChan := make(chan bool, 1)
		go func() {
			server.GracefulStop()
			gracefulStopChan <- true
		}()
		t := time.NewTimer(10 * time.Second)
		select {
		case <-gracefulStopChan:
			log.Print("Success graceful stop")
		case <-t.C:
			server.Stop()
		}
	}()

	errors := make(chan error)
	go func() {
		errors <- server.Serve(lis)
	}()

	if err := <-errors; err != nil {
		log.Fatal("Failed to server gRPC server", err)
	}

}

func RegisterHeathCheck(s *grpc.Server) {
	health.RegisterHealthServer(s, &healthServer{})
}

type healthServer struct{}

func (h *healthServer) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{
		Status: health.HealthCheckResponse_SERVING,
	}, nil
}

func (h *healthServer) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "service watch is not implemented current version.")
}

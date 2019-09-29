package server

import (
	"net"
	"net/http"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/tommy-sho/opencensus-tracing/user_service/config"
	"go.uber.org/zap"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const appName = "userService"

type Server struct {
	grpcServer     *grpc.Server
	grpcListener   net.Listener
	healthListener net.Listener
	logger         *zap.Logger
}

func Exec(cfg *config.Config) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	return nil
}

func New(appPort, healthPort string, logger *zap.Logger) (*Server, error) {
	grpcListner, err := net.Listen("tcp", appPort)
	if err != nil {
		return nil, xerrors.Errorf("can't listen app listener [port: %s]: %w", appPort, err)
	}

	healthLister, err := net.Listen("tcp", healthPort)
	if err != nil {
		return nil, xerrors.Errorf("can't listen app health [port: %s]: %w", appPort, err)
	}
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/health", func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	grpc_logrus.ReplaceGrpcLoggerV2(logger)

	grpcServer := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time: 150 * time.Second,
		}))

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(grpcServer)

	return &Server{
		grpcServer:     grpcServer,
		grpcListener:   grpcListner,
		healthListener: healthLister,
	}, nil
}

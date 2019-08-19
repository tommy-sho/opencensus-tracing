package main

import (
	"context"
	"github.com/tommy-sho/opencensus-tracing/proto"
)

const (
	port = "50001"
)

func main() {
	//b := &BackendServer{}
	//server := grpc.NewServer()
	//proto.RegisterBackendServerServer(server, b)
	//lib.RegisterHeathCheck(server)
	//reflection.Register(server)
	//lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	//if err != nil {
	//	panic(err)
	//}
	//
	//stopChan := make(chan os.Signal, 1)
	//signal.Notify(stopChan,
	//	os.Interrupt,
	//	syscall.SIGINT,
	//	syscall.SIGTERM,
	//)
	//go func() {
	//	<-stopChan
	//	gracefulStopChan := make(chan bool, 1)
	//	go func() {
	//		server.GracefulStop()
	//		gracefulStopChan <- true
	//	}()
	//	t := time.NewTimer(10 * time.Second)
	//	select {
	//	case <-gracefulStopChan:
	//		log.Print("Success graceful stop")
	//	case <-t.C:
	//		server.Stop()
	//	}
	//}()
	//
	//errors := make(chan error)
	//go func() {
	//	errors <- server.Serve(lis)
	//}()
	//
	//if err := <-errors; err != nil {
	//	log.Fatal("Failed to server gRPC server", err)
	//}

}

type UserService struct {}

func (u UserService) SaveUser(context.Context, *proto.GetUserRequest) (*proto.SaveUserResponse, error) {
	panic("implement me")
}

func (u UserService) GetUser(context.Context, *proto.GetUserRequest) (*proto.GetUserResponse, error) {
	panic("implement me")
}

func (u UserService) GetUser2(context.Context, *proto.GetUsersRequest) (*proto.GetUsersResponse, error) {
	panic("implement me")
}


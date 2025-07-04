package transportgrpc

import (
	"log"
	"net"

	taskpb "github.com/IbadT/project-protos/proto/task"
	userpb "github.com/IbadT/project-protos/proto/user"
	"github.com/IbadT/task_service-golang_microservice/internal/task"
	"google.golang.org/grpc"
)

func RunGRPC(svc task.Service, uc userpb.UserServiceClient) error {
	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Ошибка при запуске сервиса %v", err)
		return err
	}
	grpcSrv := grpc.NewServer()

	taskpb.RegisterTaskServiceServer(grpcSrv, NewHandler(svc, uc))

	log.Printf("GRPC запущен на порту 50052")
	if err := grpcSrv.Serve(listener); err != nil {
		log.Fatalf("Ошибка при запуске grpc сервера %v", err)
		return err
	}
	return nil
}

package main

import (
	"log"

	"github.com/IbadT/task_service-golang_microservice/internal/database"
	"github.com/IbadT/task_service-golang_microservice/internal/task"
	transportgrpc "github.com/IbadT/task_service-golang_microservice/internal/transport/grpc"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	repository := task.NewRepository(db)
	service := task.NewService(repository)

	// Connect to user_service by container name
	userClient, conn, err := transportgrpc.NewUserClient("user_service:50051")
	if err != nil {
		log.Fatalf("failed to connect to users: %v", err)
	}
	defer conn.Close()

	if err := transportgrpc.RunGRPC(service, userClient); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}

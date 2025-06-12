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

	if err := transportgrpc.RunGRPC(service); err != nil {
		log.Fatalf("gRPC сервер завершился с ошибкой: %v", err)
	}
}

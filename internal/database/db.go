package database

import (
	"log"

	"github.com/IbadT/task_service-golang_microservice/internal/task"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=task_db user=postgres password=postgres dbname=t_mic sslmode=disable"

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
		return nil, err
	}
	if err := DB.AutoMigrate(&task.Task{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
		return nil, err
	}
	return DB, nil
}

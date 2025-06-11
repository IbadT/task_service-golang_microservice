package task

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// go get github.com/IbadT/project-protos@latest
type Task struct {
	gorm.Model
	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title string    `json:"title" gorm:"type:varchar(255);not null"`
}

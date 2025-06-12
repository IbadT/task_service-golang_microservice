package task

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// go get github.com/IbadT/project-protos@latest
// go get github.com/IbadT/project-protos@1.0.0
// go get github.com/your-org/project-protos@v1.0.0

// type User struct {
// 	ID    uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
// 	Email string    `json:"email" gorm:"type:varchar(255);not null"`
// }

type Task struct {
	gorm.Model
	ID     uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title  string    `json:"title" gorm:"type:varchar(255);not null"`
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid"`
}

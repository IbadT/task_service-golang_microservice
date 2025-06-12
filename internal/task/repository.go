package task

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	DB *gorm.DB
}

type Repository interface {
	CreateTask(task Task) error
	GetTaskById(taskId uuid.UUID) (Task, error)
	GetTask(taskId uuid.UUID) (Task, error)
	ListTasks() ([]Task, error)
	UpdateTask(task Task) error
	DeleteTask(taskId uuid.UUID) error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) GetTaskById(taskId uuid.UUID) (Task, error) {
	var task Task
	err := r.DB.First(&task, "id = ?", taskId).Error
	return task, err
}

func (r *repository) CreateTask(task Task) error {
	return r.DB.Create(&task).Error
}

func (r *repository) GetTask(taskId uuid.UUID) (Task, error) {
	var task Task
	err := r.DB.First(&task, "id = ?", taskId).Error
	return task, err
}

func (r *repository) ListTasks() ([]Task, error) {
	var tasks []Task
	err := r.DB.Find(&tasks).Error
	return tasks, err
}

func (r *repository) UpdateTask(task Task) error {
	return r.DB.Save(&task).Error
}

func (r *repository) DeleteTask(taskId uuid.UUID) error {
	return r.DB.Delete(&Task{}, taskId).Error
}

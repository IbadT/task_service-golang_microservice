package task

import "github.com/google/uuid"

type service struct {
	repo Repository
}

type Service interface {
	CreateTask(title string, userId uuid.UUID) (Task, error)
	GetTask(taskId uuid.UUID) (Task, error)
	ListTasks() ([]Task, error)
	ListTasksByUser(userId uuid.UUID) ([]Task, error)
	UpdateTask(taskId uuid.UUID, title string) (Task, error)
	DeleteTask(taskId uuid.UUID) error
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) ListTasksByUser(userId uuid.UUID) ([]Task, error) {
	return s.repo.ListTasksByUser(userId)
}

func (s *service) CreateTask(title string, userId uuid.UUID) (Task, error) {
	task := Task{
		ID:     uuid.New(),
		Title:  title,
		UserID: userId,
	}

	err := s.repo.CreateTask(task)
	if err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *service) GetTask(taskId uuid.UUID) (Task, error) {
	return s.repo.GetTask(taskId)
}

func (s *service) ListTasks() ([]Task, error) {
	return s.repo.ListTasks()
}

func (s *service) UpdateTask(taskId uuid.UUID, title string) (Task, error) {
	task, err := s.repo.GetTaskById(taskId)
	if err != nil {
		return Task{}, err
	}

	task.Title = title
	if err := s.repo.UpdateTask(task); err != nil {
		return Task{}, err
	}
	return task, nil
}

func (s *service) DeleteTask(taskId uuid.UUID) error {
	return s.repo.DeleteTask(taskId)
}

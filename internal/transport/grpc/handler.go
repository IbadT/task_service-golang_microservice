package transportgrpc

import (
	"context"
	"fmt"

	taskpb "github.com/IbadT/project-protos/proto/task"
	userpb "github.com/IbadT/project-protos/proto/user"
	"github.com/IbadT/task_service-golang_microservice/internal/task"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc        task.Service
	userClient userpb.UserServiceClient
	taskpb.UnimplementedTaskServiceServer
}

func NewHandler(s task.Service, uc userpb.UserServiceClient) *Handler {
	return &Handler{
		svc:        s,
		userClient: uc,
	}
}

func (h *Handler) ListTasksByUser(ctx context.Context, req *taskpb.ListTasksByUserRequest) (*taskpb.ListTasksByUserResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return &taskpb.ListTasksByUserResponse{}, err
	}
	tasks, err := h.svc.ListTasksByUser(userId)
	if err != nil {
		return &taskpb.ListTasksByUserResponse{}, err
	}

	pbTasks := make([]*taskpb.Task, 0, len(tasks))
	for _, t := range tasks {
		tsk := &taskpb.Task{
			Id:     t.ID.String(),
			Title:  t.Title,
			UserId: t.UserID.String(),
		}
		pbTasks = append(pbTasks, tsk)
	}
	return &taskpb.ListTasksByUserResponse{Task: pbTasks}, nil
}

func (h *Handler) CreateTask(ctx context.Context, req *taskpb.CreateTaskRequest) (*taskpb.CreateTaskResponse, error) {
	pbUser := &userpb.GetUserRequest{
		Id: req.UserId,
	}
	if _, err := h.userClient.GetUser(ctx, pbUser); err != nil {
		return nil, fmt.Errorf("user %s not found: %w", req.UserId, err)
	}

	title := req.Title
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return &taskpb.CreateTaskResponse{}, err
	}

	task, err := h.svc.CreateTask(title, userId)
	if err != nil {
		return &taskpb.CreateTaskResponse{}, err
	}

	pbTask := &taskpb.Task{
		Id:     task.ID.String(),
		Title:  task.Title,
		UserId: task.UserID.String(),
	}
	return &taskpb.CreateTaskResponse{Task: pbTask}, nil
}

func (h *Handler) GetTask(ctx context.Context, req *taskpb.GetTaskRequest) (*taskpb.GetTaskResponse, error) {
	taskId, err := uuid.Parse(req.Id)
	if err != nil {
		return &taskpb.GetTaskResponse{}, err
	}

	task, err := h.svc.GetTask(taskId)
	if err != nil {
		return &taskpb.GetTaskResponse{}, err
	}

	pbTask := &taskpb.Task{
		Id:     task.ID.String(),
		Title:  task.Title,
		UserId: task.UserID.String(),
	}
	return &taskpb.GetTaskResponse{Task: pbTask}, nil
}

func (h *Handler) ListTasks(ctx context.Context, req *emptypb.Empty) (*taskpb.ListTasksResponse, error) {
	tasks, err := h.svc.ListTasks()
	if err != nil {
		return &taskpb.ListTasksResponse{}, err
	}
	pbTasks := make([]*taskpb.Task, 0, len(tasks))
	for _, t := range tasks {
		tsk := &taskpb.Task{
			Id:     t.ID.String(),
			Title:  t.Title,
			UserId: t.UserID.String(),
		}
		pbTasks = append(pbTasks, tsk)
	}
	return &taskpb.ListTasksResponse{Task: pbTasks}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *taskpb.UpdateTaskRequest) (*taskpb.UpdateTaskResponse, error) {
	title := req.Title
	taskId, err := uuid.Parse(req.Id)
	if err != nil {
		return &taskpb.UpdateTaskResponse{}, err
	}

	task, err := h.svc.UpdateTask(taskId, title)
	if err != nil {
		return &taskpb.UpdateTaskResponse{}, err
	}
	pbTask := &taskpb.Task{
		Id:     task.ID.String(),
		Title:  task.Title,
		UserId: task.UserID.String(),
	}
	return &taskpb.UpdateTaskResponse{Task: pbTask}, nil

}

func (h *Handler) DeleteTask(ctx context.Context, req *taskpb.DeleteTaskRequest) (*taskpb.DeleteTaskResponse, error) {
	taskId, err := uuid.Parse(req.Id)
	if err != nil {
		return &taskpb.DeleteTaskResponse{}, err
	}

	if err := h.svc.DeleteTask(taskId); err != nil {
		return &taskpb.DeleteTaskResponse{}, err
	}

	return &taskpb.DeleteTaskResponse{Id: req.Id}, nil
}

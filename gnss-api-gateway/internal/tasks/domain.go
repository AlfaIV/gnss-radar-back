package tasks_domain_gateway

import (
	"context"
	user_domain_gateway "gnss-radar/gnss-api-gateway/internal/user"
)

type Task struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	DateTimeStart string   `json:"datetimeStart"`
	DateTimeEnd   string   `json:"datetimeEnd"`
	CreatorId     string   `json:"creatorId"`
	Satellites    []string `json:"satellites"`
}

type GetTaskResponseEntity struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	DateTimeStart string   `json:"datetimeStart"`
	DateTimeEnd   string   `json:"datetimeEnd"`
	CreatorId     string   `json:"creatorId"`
	Satellites    []string `json:"satellites"`
	user_domain_gateway.UserData
}

type GetTasksResponse struct {
	Tasks []GetTaskResponseEntity `json:"tasks"`
}

type DeleteTaskRequest struct {
	Id string `json:"id"`
}

type Usecase interface {
	CreateTask(ctx context.Context, task Task) error
	GetTasks(ctx context.Context, size uint64, page uint64) ([]Task, error)
	UpdateTask(ctx context.Context, task Task) error
	DeleteTask(ctx context.Context, id string) error
}

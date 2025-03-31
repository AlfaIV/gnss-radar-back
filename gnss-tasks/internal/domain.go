package tasks_domain

import "context"

type Task struct {
	Id            string
	Name          string
	Description   string
	DateTimeStart string
	DateTimeEnd   string
	CreatorId     string
	IsAll         bool
	Satellites    []string
}

type Repository interface {
	CreateTask(ctx context.Context, r Task) error
	GetTasks(ctx context.Context, size uint64, page uint64) ([]Task, error)
	UpdateTask(ctx context.Context, r Task) error
	DeleteTask(ctx context.Context, id string) error
}

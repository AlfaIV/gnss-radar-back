package auth_client

import (
	"context"
	common_proto "gnss-radar/api/proto/common"
	proto "gnss-radar/api/proto/tasks"
	tasks_domain_gateway "gnss-radar/gnss-api-gateway/internal/tasks"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type TasksClient struct {
	client proto.TasksClient
	logger *logrus.Logger
}

func NewTasksClient(client proto.TasksClient, logger *logrus.Logger) TasksClient {
	return TasksClient{client: client, logger: logger}
}

func (tc *TasksClient) CreateTask(ctx context.Context, task tasks_domain_gateway.Task) error {
	_, err := tc.client.CreateTask(ctx, &proto.Task{
		Name: task.Name,
		Description: task.Description,
		DateTimeStart: task.DateTimeStart,
		DateTimeEnd: task.DateTimeEnd,
		CreatorId: task.CreatorId,
		IsAll: task.IsAll,
		Satellites: task.Satellites,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create task")
	}

	return nil
}

func (tc *TasksClient) GetTasks(ctx context.Context, size uint64, page uint64) ([]tasks_domain_gateway.Task, error) {
	tasksArray, err := tc.client.GetTasks(ctx, &common_proto.PaginatedRequest{Size: size, Page: page})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get tasks")
	}

	var tasks []tasks_domain_gateway.Task

	for _, task := range(tasksArray.GetTasks()) {
		tasks = append(tasks, tasks_domain_gateway.Task{
			Id: task.GetId(),
			Name: task.GetName(),
			Description: task.GetDescription(),
			DateTimeStart: task.GetDateTimeStart(),
			DateTimeEnd: task.GetDateTimeEnd(),
			CreatorId: task.GetCreatorId(),
			Satellites: task.GetSatellites(),
		})
	}

	return tasks, nil
}

func (tc *TasksClient) UpdateTask(ctx context.Context, task tasks_domain_gateway.Task) error {
	if _, err := tc.client.UpdateTask(ctx, &proto.Task{
		Name: task.Name,
		Description: task.Description,
		DateTimeStart: task.DateTimeStart,
		DateTimeEnd: task.DateTimeEnd,
		CreatorId: task.CreatorId,
		IsAll: task.IsAll,
		Satellites: task.Satellites,
	}); err != nil {
		return errors.Wrap(err, "failed to update task")
	}

	return nil
}

func (tc *TasksClient) DeleteTask(ctx context.Context, id string) error {
	_, err := tc.client.DeleteTask(ctx, &proto.TaskId{TaskId: id})
	if err != nil {
		return errors.Wrapf(err, "failed to get delete task")
	}

	return nil
}

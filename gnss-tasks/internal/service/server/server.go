package tasks_server

import (
	"context"
	common_proto "gnss-radar/api/proto/common"
	proto "gnss-radar/api/proto/tasks"

	tasks_domain "gnss-radar/gnss-tasks/internal"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type TasksServiceServer struct {
	logger *logrus.Logger
	repo   tasks_domain.Repository

	proto.UnimplementedTasksServer
}

func NewTasksServer(repo tasks_domain.Repository, logger *logrus.Logger) TasksServiceServer {
	return TasksServiceServer{repo: repo, logger: logger}
}

func (s *TasksServiceServer) CreateTask(ctx context.Context, r *proto.Task) (*emptypb.Empty, error) {

	err := s.repo.CreateTask(ctx, tasks_domain.Task{
		Name:          r.GetName(),
		Description:   r.GetDescription(),
		DateTimeStart: r.GetDateTimeStart(),
		DateTimeEnd:   r.GetDateTimeEnd(),
		CreatorId:     r.GetCreatorId(),
		IsAll:         r.GetIsAll(),
		Satellites:    r.GetSatellites(),
	})

	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "[TASKS]: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *TasksServiceServer) GetTasks(ctx context.Context, r *common_proto.PaginatedRequest) (*proto.TasksArray, error) {

	tasks, err := s.repo.GetTasks(ctx, r.GetSize(), r.GetPage())
	if err != nil {
		return &proto.TasksArray{}, status.Errorf(codes.Internal, "[TASKS]: %v", err)
	}

	var taskList []*proto.Task
	for _, task := range tasks {
		taskList = append(taskList, &proto.Task{
			Id:            task.Id,
			Name:          task.Name,
			Description:   task.Description,
			DateTimeStart: task.DateTimeStart,
			DateTimeEnd:   task.DateTimeEnd,
			CreatorId:     task.CreatorId,
			Satellites:    task.Satellites,
			IsAll:         task.IsAll,
		})
	}

	return &proto.TasksArray{Tasks: taskList}, nil
}

func (s *TasksServiceServer) UpdateTask(ctx context.Context, r *proto.Task) (*emptypb.Empty, error) {

	err := s.repo.UpdateTask(ctx, tasks_domain.Task{
		Id:            r.GetId(),
		Name:          r.GetName(),
		Description:   r.GetDescription(),
		DateTimeStart: r.GetDateTimeStart(),
		DateTimeEnd:   r.GetDateTimeEnd(),
		IsAll:         r.GetIsAll(),
		Satellites:    r.GetSatellites(),
	})
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "[TASKS]: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *TasksServiceServer) DeleteTask(ctx context.Context, r *proto.TaskId) (*emptypb.Empty, error) {

	err := s.repo.DeleteTask(ctx, r.GetTaskId())
	if err != nil {
		return &emptypb.Empty{}, status.Errorf(codes.Internal, "[TASKS]: %v", err)
	}

	return &emptypb.Empty{}, nil
}

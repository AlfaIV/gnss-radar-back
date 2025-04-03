package tasks_handler

import (
	"fmt"
	tasks_domain_gateway "gnss-radar/gnss-api-gateway/internal/tasks"
	user_domain_gateway "gnss-radar/gnss-api-gateway/internal/user"
	"gnss-radar/gnss-api-gateway/pkg/mwutils"
	"net/http"
	"strconv"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type TasksHandler struct {
	userUsecase user_domain_gateway.Usecase
	taskUsecase tasks_domain_gateway.Usecase
	logger      *logrus.Logger
}

func NewHandler(
	user user_domain_gateway.Usecase,
	tasks tasks_domain_gateway.Usecase,
	logger *logrus.Logger,
) TasksHandler {
	return TasksHandler{
		userUsecase: user,
		taskUsecase: tasks,
		logger:      logger,
	}
}

func (h *TasksHandler) CreateTask(c echo.Context) error {

	ctx := c.Request().Context()

	id, err := mwutils.GetUserID(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "No session id provided")
	}

	task := tasks_domain_gateway.Task{}
	if err := c.Bind(&task); err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusBadRequest, "failed to parse request data")
	}

	task.CreatorId = id

	err = h.taskUsecase.CreateTask(c.Request().Context(), task)
	if err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusInternalServerError, "an arror occured")
	}

	return c.NoContent(http.StatusOK)
}

func (h *TasksHandler) GetTasks(c echo.Context) error {

	ctx := c.Request().Context()

	_, err := mwutils.GetUserID(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "No session id provided")
	}

	pageParam := c.QueryParam("page")
	if pageParam == "" {

		return c.String(http.StatusBadRequest, "Incorrect page param")
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusBadRequest, "Incorrect page param")
	}

	sizeParam := c.QueryParam("size")
	if pageParam == "" {

		return c.String(http.StatusBadRequest, "Incorrect size param")
	}

	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusBadRequest, "Incorrect size param")
	}

	tasks, err := h.taskUsecase.GetTasks(ctx, uint64(size), uint64(page))
	if err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusInternalServerError, "failed to get tasks")
	}

	tasksPtr := make([]*tasks_domain_gateway.GetTaskResponseEntity, len(tasks))
	for i := range tasks {
		tasksPtr[i] = &tasks_domain_gateway.GetTaskResponseEntity{
			Id:            tasks[i].Id,
			Name:          tasks[i].Name,
			Description:   tasks[i].Description,
			DateTimeStart: tasks[i].DateTimeStart,
			DateTimeEnd:   tasks[i].DateTimeEnd,
			CreatorId:     tasks[i].CreatorId,
			Satellites:    tasks[i].Satellites,
		}
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(tasks))
	sem := make(chan struct{}, 10)

	for idx := range tasksPtr {
		wg.Add(1)
		sem <- struct{}{}

		go func(i int) {
			defer func() {
				<-sem
				wg.Done()
			}()

			data, err := h.userUsecase.GetUserInfoById(ctx, tasksPtr[i].CreatorId)
			if err != nil {
				errChan <- fmt.Errorf("task %d: %w", i, err)
				return
			}

			tasksPtr[i].UserName = data.Name
			tasksPtr[i].Surname = data.Surname
			tasksPtr[i].Email = data.Email
			tasksPtr[i].OrganizationName = data.OrganizationName
		}(idx)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for err := range errChan {
		h.logger.Error("[GW]: ", err)
	}

	responseTasks := make([]tasks_domain_gateway.GetTaskResponseEntity, len(tasksPtr))
	for i, taskPtr := range tasksPtr {
		responseTasks[i] = *taskPtr
	}

	return c.JSON(http.StatusOK, tasks_domain_gateway.GetTasksResponse{
		Tasks: responseTasks,
	})
}

func (h *TasksHandler) UpdateTask(c echo.Context) error {
	ctx := c.Request().Context()

	_, err := mwutils.GetUserID(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "No session id provided")
	}

	task := tasks_domain_gateway.Task{}
	if err := c.Bind(&task); err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusBadRequest, "failed to parse request data")
	}

	if err := h.taskUsecase.UpdateTask(c.Request().Context(), task); err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusBadRequest, "failed to update task")
	}

	return c.NoContent(http.StatusOK)
}

func (h *TasksHandler) DeleteTask(c echo.Context) error {
	ctx := c.Request().Context()

	_, err := mwutils.GetUserID(ctx)
	if err != nil {
		h.logger.Error("[GW]: ", err)

		return c.String(http.StatusUnauthorized, "No session id provided")
	}

	delTask := tasks_domain_gateway.DeleteTaskRequest{}
	if err := c.Bind(&delTask); err != nil {
		h.logger.Error("[GW]:", err)
		return c.String(http.StatusBadRequest, "failed to parse request data")
	}

	if err := h.taskUsecase.DeleteTask(c.Request().Context(), delTask.Id); err != nil {
		h.logger.Error("[GW]: ", err)
		return c.String(http.StatusBadRequest, "failed to delete task")
	}

	return c.NoContent(http.StatusOK)
}

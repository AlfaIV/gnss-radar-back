package tasks_domain_gateway

import (
	"context"
	measurements_domain_gateway "gnss-radar/gnss-api-gateway/internal/measurements"
	user_domain_gateway "gnss-radar/gnss-api-gateway/internal/user"
	"time"
)

type Task struct {
	Id            string   `json:"id"`
	Name          string   `json:"name"`
	Description   string   `json:"description"`
	DateTimeStart string   `json:"datetimeStart"`
	DateTimeEnd   string   `json:"datetimeEnd"`
	CreatorId     string   `json:"creatorId"`
	IsAll         bool     `json:"isAll"`
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
	IsAll         bool     `json:"isAll"`
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

// TAH4UK
func ValidateSatellitesIntersection(
	satellitesData measurements_domain_gateway.SatelliteWithIntervals,
	endDatetime string,
	startDatetime string,
	isAll bool) bool {

	const intervalConst = 0 // Отсечка

	if (isAll == true) && (satellitesData.Intervals == nil) {
		return true
	}
	// Реализация без обработок ошибок
	endDateParsed, _ := time.Parse(time.RFC3339, endDatetime+"Z")
	startDateParsed, _ := time.Parse(time.RFC3339, startDatetime+"Z")

	if startDateParsed.Before(endDateParsed) {
		for _, value := range satellitesData.Intervals {
			endSat, _ := time.Parse(time.RFC3339, value.EndDatetime+"Z")
			startSat, _ := time.Parse(time.RFC3339, value.StartDatetime+"Z")
			if startSat.Before(startDateParsed) || endSat.After(endDateParsed) ||
				endSat.Sub(startSat).Minutes() < intervalConst {
				return false
			}
		}
		return true
	}
	return false
}

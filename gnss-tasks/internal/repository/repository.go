package tasks_repository

import (
	"context"
	tasks_domain "gnss-radar/gnss-tasks/internal"
	"strings"
	"sync"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type PgxIFace interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

type TaskRepo struct {
	pool   PgxIFace
	logger *logrus.Logger
}

func NewTaskRepo(pool PgxIFace, logger *logrus.Logger) *TaskRepo {
	return &TaskRepo{pool: pool, logger: logger}
}

type processedSatellite struct {
	name string
}

func satelliteWorker(ctx context.Context, wg *sync.WaitGroup, in <-chan string, out chan<- processedSatellite) {
	defer wg.Done()
	for name := range in {
		select {
		case <-ctx.Done():
			return
		default:
			out <- processedSatellite{
				name: strings.TrimSpace(name),
			}
		}
	}
}

func (tr *TaskRepo) CreateTask(ctx context.Context, r tasks_domain.Task) error {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return errors.Wrap(err, "failed to load Moscow location")
	}

	parseTime := func(timeStr string) (time.Time, error) {
		return time.ParseInLocation("2006-01-02T15:04:05", timeStr, loc)
	}

	startTime, err := parseTime(r.DateTimeStart)
	if err != nil {
		return errors.Wrap(err, "invalid start time format")
	}

	endTime, err := parseTime(r.DateTimeEnd)
	if err != nil {
		return errors.Wrap(err, "invalid end time format")
	}

	now := time.Now().In(loc)
	if startTime.After(endTime) {
		return errors.New("start time cannot be after end time")
	}

	if endTime.Before(now) {
		return errors.New("end time cannot be in the past")
	}

	tx, err := tr.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	defer tx.Rollback(ctx)

	var taskId string
	err = tx.QueryRow(
		ctx,
		`INSERT INTO task(name, description, time_start, time_end, creator_id, is_all)
         VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`,
		r.Name,
		r.Description,
		startTime,
		endTime,
		r.CreatorId,
		r.IsAll,
	).Scan(&taskId)

	if err != nil {
		tr.logger.WithError(err).Error("task repo: create failed")
		return errors.Wrap(err, "failed to create task")
	}

	if !r.IsAll && len(r.Satellites) > 0 {
		inputChan := make(chan string, len(r.Satellites))
		resultChan := make(chan processedSatellite)

		var wg sync.WaitGroup
		const workers = 4

		for i := 0; i < workers; i++ {
			wg.Add(1)
			go satelliteWorker(ctx, &wg, inputChan, resultChan)
		}

		go func() {
			wg.Wait()
			close(resultChan)
		}()

		go func() {
			defer close(inputChan)
			for _, name := range r.Satellites {
				select {
				case inputChan <- name:
				case <-ctx.Done():
					return
				}
			}
		}()

		unique := make(map[string]struct{})

		for res := range resultChan {
			if res.name != "" {
				unique[res.name] = struct{}{}
			}
		}

		var satellites []string
		for name := range unique {
			satellites = append(satellites, name)
		}

		if len(satellites) > 0 {
			batch := &pgx.Batch{}
			for _, name := range satellites {
				batch.Queue(
					`INSERT INTO task_satellites(task_id, satellite_name) VALUES ($1, $2)`,
					taskId,
					name,
				)
			}

			br := tx.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				return errors.Wrap(err, "failed to insert satellites")
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

// Потом ещё фильтр
func (tr *TaskRepo) GetTasks(ctx context.Context, size uint64, page uint64) ([]tasks_domain.Task, error) {
	query := `
        SELECT
			t.id::text,
            t.name,
            t.description,
			t.is_all,
            TO_CHAR(t.time_start AT TIME ZONE 'Europe/Moscow', 'YYYY-MM-DD"T"HH24:MI:SS'),
            TO_CHAR(t.time_end AT TIME ZONE 'Europe/Moscow', 'YYYY-MM-DD"T"HH24:MI:SS'),
            t.creator_id::text,
            COALESCE(array_agg(ts.satellite_name) FILTER (WHERE ts.satellite_name IS NOT NULL), '{}') AS satellites
        FROM task t
        LEFT JOIN task_satellites ts ON t.id = ts.task_id
        GROUP BY t.id
        ORDER BY t.time_start DESC
        LIMIT $1 OFFSET $2
    `

	rows, err := tr.pool.Query(ctx, query, size, page-1)
	if err != nil {
		return nil, errors.Wrap(err, "failed to query tasks")
	}
	defer rows.Close()

	var tasks []tasks_domain.Task
	for rows.Next() {
		var task tasks_domain.Task
		var satellites []string

		err := rows.Scan(
			&task.Id,
			&task.Name,
			&task.Description,
			&task.IsAll,
			&task.DateTimeStart,
			&task.DateTimeEnd,
			&task.CreatorId,
			&satellites,
		)

		if err != nil {
			return nil, errors.Wrap(err, "failed to scan task row")
		}

		task.Satellites = satellites
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows iteration error")
	}

	return tasks, nil
}

func (tr *TaskRepo) UpdateTask(ctx context.Context, r tasks_domain.Task) error {
	loc, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		return errors.Wrap(err, "failed to load Moscow location")
	}

	parseTime := func(timeStr string) (time.Time, error) {
		return time.ParseInLocation("2006-01-02T15:04:05", timeStr, loc)
	}

	startTime, err := parseTime(r.DateTimeStart)
	if err != nil {
		return errors.Wrap(err, "invalid start time format")
	}

	endTime, err := parseTime(r.DateTimeEnd)
	if err != nil {
		return errors.Wrap(err, "invalid end time format")
	}

	now := time.Now().In(loc)
	if startTime.After(endTime) {
		return errors.New("start time cannot be after end time")
	}

	if endTime.Before(now) {
		return errors.New("end time cannot be in the past")
	}

	tx, err := tr.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	defer tx.Rollback(ctx)

	updateTaskQuery := `
        UPDATE task 
        SET 
            name = $1,
            description = $2,
            time_start = $3,
            time_end = $4,
			is_all = $5
        WHERE id = $6
    `

	_, err = tx.Exec(ctx,
		updateTaskQuery,
		r.Name,
		r.Description,
		startTime,
		endTime,
		r.IsAll,
		r.Id,
	)

	if err != nil {
		return errors.Wrap(err, "failed to update task")
	}

	if _, err := tx.Exec(ctx, "DELETE FROM task_satellites WHERE task_id = $1", r.Id); err != nil {
		return errors.Wrap(err, "failed to delete old satellites")
	}

	if !r.IsAll && len(r.Satellites) > 0 {

		inputChan := make(chan string, len(r.Satellites))
		resultChan := make(chan processedSatellite)

		var wg sync.WaitGroup
		const workers = 4

		for i := 0; i < workers; i++ {
			wg.Add(1)
			go satelliteWorker(ctx, &wg, inputChan, resultChan)
		}

		go func() {
			wg.Wait()
			close(resultChan)
		}()

		go func() {
			defer close(inputChan)
			for _, name := range r.Satellites {
				select {
				case inputChan <- name:
				case <-ctx.Done():
					return
				}
			}
		}()

		unique := make(map[string]struct{})

	loop:
		for {
			select {
			case res, ok := <-resultChan:
				if !ok {
					break loop
				}
				unique[res.name] = struct{}{}
			case <-ctx.Done():
				return ctx.Err()
			}
		}

		var satellitesToInsert []string
		for name := range unique {
			satellitesToInsert = append(satellitesToInsert, name)
		}

		if len(satellitesToInsert) > 0 {
			batch := &pgx.Batch{}
			for _, name := range satellitesToInsert {
				batch.Queue(
					"INSERT INTO task_satellites(task_id, satellite_name) VALUES ($1, $2)",
					r.Id,
					name,
				)
			}

			br := tx.SendBatch(ctx, batch)
			if err := br.Close(); err != nil {
				return errors.Wrap(err, "failed to insert new satellites")
			}
		}
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil
}

func (tr *TaskRepo) DeleteTask(ctx context.Context, id string) error {
	deleteTaskSatellitesQuery := `
		DELETE FROM task_satellites WHERE task_id=$1;
	`

	deleteTaskQuery := `
		DELETE FROM task WHERE id = $1;
	`
	tx, err := tr.pool.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to start transaction")
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Query(ctx, deleteTaskSatellitesQuery, id); err != nil {
		return errors.Wrap(err, "Failed to delete satellites for task")
	}

	if _, err = tx.Query(ctx, deleteTaskQuery, id); err != nil {
		return errors.Wrap(err, "Failed to delete task")
	}

	if err := tx.Commit(ctx); err != nil {
		return errors.Wrap(err, "failed to commit transaction")
	}

	return nil

}

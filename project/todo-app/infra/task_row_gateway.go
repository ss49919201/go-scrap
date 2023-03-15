package infra

import (
	"database/sql"
	"time"
	tablerowgateway "todo-app/table_row_gateway"
)

type taskRowGateway struct {
	ID          int
	Name        string
	Description string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *taskRowGateway) Insert() error

func (t *taskRowGateway) Update() error

func (t *taskRowGateway) Delete() error

func NewTaskFinder(db *sql.DB) tablerowgateway.TaskFinder {
	return &taskFinder{db}
}

type taskFinder struct {
	*sql.DB
}

func (t *taskFinder) First(id int) (tablerowgateway.TaskRowGateway, error)

func (t *taskFinder) FindByUserID(userID int) ([]tablerowgateway.TaskRowGateway, error)

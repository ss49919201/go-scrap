package infra

import (
	"database/sql"
	"time"
)

type TaskRowGateway struct {
	ID          int
	Name        string
	Description string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *TaskRowGateway) Insert() error

func (t *TaskRowGateway) Update() error

func (t *TaskRowGateway) Delete() error

type TaskFinder interface {
	First(id int) (*TaskRowGateway, error)
	FindByUserID(userID int) ([]*TaskRowGateway, error)
}

func NewTaskFinder(db *sql.DB) TaskFinder {
	return &taskFinder{db}
}

type taskFinder struct {
	*sql.DB
}

func (t *taskFinder) First(id int) (*TaskRowGateway, error)

func (t *taskFinder) FindByUserID(userID int) ([]*TaskRowGateway, error)

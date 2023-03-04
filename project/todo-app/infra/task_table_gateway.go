package infra

import (
	"database/sql"
	"time"
)

type TaskRecord struct {
	ID          int
	Name        string
	Description string
	UserID      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskTableGateway interface {
	First(id int) (*TaskRecord, error)
	FindByUserID(userID int) ([]*TaskRecord, error)
	Insert(name, description string, userID int) (*TaskRecord, error)
	Update(id int, name, description string, userID int) (*TaskRecord, error)
	Delete(id int) error
}

func NewTaskTableGateway(db *sql.DB) TaskTableGateway {
	return &taskTableGateway{db}
}

type taskTableGateway struct {
	*sql.DB
}

func (t *taskTableGateway) First(id int) (*TaskRecord, error)

func (t *taskTableGateway) FindByUserID(userID int) ([]*TaskRecord, error)

func (t *taskTableGateway) Insert(name, description string, userID int) (*TaskRecord, error)

func (t *taskTableGateway) Update(id int, name, description string, userID int) (*TaskRecord, error)

func (t *taskTableGateway) Delete(id int) error

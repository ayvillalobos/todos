package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Task model struct.
type Task struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Title       string    `json:"title" db:"title"`
	LimitDate   time.Time `json:"limit_date" db:"limit_date"`
	Description string    `json:"description" db:"-"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Tasks array model struct of Task.
type Tasks []Task

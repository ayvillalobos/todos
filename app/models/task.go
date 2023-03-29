package models

import (
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Task model struct.
type Task struct {
	ID          uuid.UUID  `json:"id" db:"id"`
	Title       string     `json:"title" db:"title"`
	LimitDate   nulls.Time `json:"limit_date" db:"limit_date"`
	Description string     `json:"description" db:"description"`
	Complete    bool       `json:"complete" db:"complete"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

// Tasks array model struct of Task.
type Tasks []Task

func (t *Task) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Title, Name: "Title"},
		&validators.TimeIsPresent{Field: t.LimitDate.Time, Name: "Limit Date"},
		&validators.StringIsPresent{Field: t.Description, Name: "Description"},
	), nil
}

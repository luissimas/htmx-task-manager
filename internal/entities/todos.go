package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `db:"id" form:"id"`
	Text      string    `db:"text" form:"text"`
	Done      bool      `db:"done" form:"done"`
	CreatedAt time.Time `db:"created_at" form:"created_at"`
}

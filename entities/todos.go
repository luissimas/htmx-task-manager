package entities

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	ID        uuid.UUID `db:"id"`
	Text      string    `db:"text"`
	Done      bool      `db:"done"`
	CreatedAt time.Time `db:"created_at"`
}

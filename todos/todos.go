package todos

import (
	"time"

	"github.com/google/uuid"
)

type Todo struct {
	Id        uuid.UUID `db:"id"`
	Text      string    `db:"text"`
	Done      bool      `db:"done"`
	CreatedAt time.Time `db:"created_at"`
}

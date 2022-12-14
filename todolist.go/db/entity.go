package db

// schema.go provides data models in DB
import (
	"database/sql"
	"time"
)

// Task corresponds to a row in `tasks` table
type Task struct {
	ID          uint64       `db:"id"`
	Title       string       `db:"title"`
	CreatedAt   time.Time    `db:"created_at"`
	Deadline    sql.NullTime `db:"deadline"`
	IsDone      bool         `db:"is_done"`
	Description string       `db:"description"`
}

type User struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	Password  []byte    `db:"password"`
	UpdatedAt time.Time `db:"updated_at"`
	CreatedAt time.Time `db:"created_at"`
	Valid     bool      `db:"valid"`
}

package db

// schema.go provides data models in DB
import (
	"time"
)

// Task corresponds to a row in `thread` table
type Thread struct {
	ID        uint64    `db:"id"`
	Title     string    `db:"title"`
	CreatedAt time.Time `db:"created_at"`
}

// User corresponds to
type User struct {
	ID         uint64    `db:"id"`
	Name       string    `db:"name"`
	Password   []byte    `db:"password"`
	Updated_at time.Time `db:"updated_at"`
}

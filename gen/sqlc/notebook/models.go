// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package notebook

import (
	"database/sql"
	"time"
)

type Entry struct {
	ID         int32
	Text       sql.NullString
	CreatorID  string
	CreatedAt  time.Time
	UpdatedAt  sql.NullTime
	DeleteTime sql.NullTime
}
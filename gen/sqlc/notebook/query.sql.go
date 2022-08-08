// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0
// source: query.sql

package notebook

import (
	"context"
	"database/sql"
	"time"
)

const getEntry = `-- name: GetEntry :one
SELECT id, text, creator_id, created_at, updated_at
FROM entries
WHERE id = $1
AND delete_time IS NULL
`

type GetEntryRow struct {
	ID        int32
	Text      sql.NullString
	CreatorID string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

func (q *Queries) GetEntry(ctx context.Context, id int32) (GetEntryRow, error) {
	row := q.db.QueryRowContext(ctx, getEntry, id)
	var i GetEntryRow
	err := row.Scan(
		&i.ID,
		&i.Text,
		&i.CreatorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

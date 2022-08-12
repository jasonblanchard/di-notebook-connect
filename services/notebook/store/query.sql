-- name: GetEntryByIdAndAuthor :one
SELECT id, text, creator_id, created_at, updated_at
FROM entries
WHERE id = $1 AND creator_id = $2
AND delete_time IS NULL;

-- name: ListEntriesByAuthor :many
SELECT
	*
FROM
	entries
WHERE
	creator_id = $1
	AND delete_time IS NULL
	AND id > $2
LIMIT $3;

-- name: CountEntriesByAuthor :one
SELECT
    count(*)
FROM entries 
WHERE
	creator_id = $1
	AND delete_time IS NULL;

-- name: CountEntriesByAuthorAfterCursor :one
SELECT
    count(*)
FROM entries 
WHERE
	creator_id = $1
	AND delete_time IS NULL
    AND id > $2;

-- name: CreateEntry :one
INSERT INTO entries (text, creator_id, created_at)
VALUES ($1, $2, $3)
RETURNING id;

-- name: UpdateEntryText :one
UPDATE entries
SET text = $1, updated_at = $2
WHERE id = $3
AND delete_time is null
RETURNING id, text, creator_id, created_at, updated_at;

-- name: DeleteEntryByIdAndAuthor :one
UPDATE entries
SET delete_time = $1
WHERE id = $2
AND creator_id = $3
RETURNING id, text, creator_id, created_at, updated_at, delete_time;

-- name: UnDeleteEntryByIdAndAuthor :one
UPDATE entries
SET delete_time = NULL
WHERE id = $1
AND creator_id = $2
RETURNING id, text, creator_id, created_at, updated_at, delete_time;
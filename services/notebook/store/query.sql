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
	entries.creator_id = $1
	AND entries.delete_time IS NULL
	AND entries.id <= CASE WHEN sqlc.arg(cursor)::int = 0 THEN
	(
		SELECT
			id
		FROM
			entries
		WHERE
			creator_id = $1
			AND delete_time IS NULL
		ORDER BY
			id DESC
		LIMIT 1)
		WHEN sqlc.arg(cursor)::int != 0 THEN sqlc.arg(cursor)::int - 1
	END
ORDER BY
	id DESC
LIMIT $2;

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
	entries.creator_id = $1
	AND delete_time IS NULL
    AND id < CASE WHEN sqlc.arg(cursor)::int = 0 THEN
	(
		SELECT
			id
		FROM
			entries
		WHERE
			creator_id = $1
			AND delete_time IS NULL
		ORDER BY
			id DESC
		LIMIT 1)
		WHEN sqlc.arg(cursor)::int != 0 THEN sqlc.arg(cursor)::int
	END;

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
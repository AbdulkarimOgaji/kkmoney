-- name: CreateEntry :exec
INSERT INTO entries (account_id, amount)
VALUES ($1, $2);

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1
LIMIT 1;

-- name: GetEntries :many
SELECT * FROM entries;

-- name: UpdateEntryAmount :exec
UPDATE entries SET amount = $2 WHERE id = $1;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = $1
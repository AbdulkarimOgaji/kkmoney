-- name: CreateTxn :exec
INSERT INTO txns (from_account_id, to_account_id, amount)
VALUES ($1, $2, $3);

-- name: GetTxn :one
SELECT * FROM txns
WHERE id = $1
LIMIT 1;

-- name: GetTxns :many
SELECT * FROM txns;

-- name: UpdateTxnAmount :exec
UPDATE txns SET amount = $2 WHERE id = $1;

-- name: DeleteTxn :exec
DELETE FROM txns
WHERE id = $1
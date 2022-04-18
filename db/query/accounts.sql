-- name: CreateAccount :exec
INSERT INTO accounts (holder, balance, currency)
VALUES ($1, $2, $3);

-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1
LIMIT 1;

-- name: GetAccounts :many
SELECT * FROM accounts;

-- name: UpdateAccountCurrency :exec
UPDATE accounts SET currency = $2 WHERE id = $1;

-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = $1

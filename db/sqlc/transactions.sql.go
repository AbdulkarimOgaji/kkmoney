// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: Txns.sql

package db

import (
	"context"

)

const createTxn = `-- name: CreateTxn :exec
INSERT INTO txns (from_account_id, to_account_id, amount)
VALUES (?, ?, ?)
`
type CreateTxnArgs struct {
	From_account_id int64 `json:"from_account_id" binding:"required"`
	To_account_id int64 `json:"to_account_id" binding:"required"`
	Amount int64 `json:"amount" binding:"required"`
}

func (q *Queries) CreateTxn(ctx context.Context, args CreateTxnArgs) (Txn, error) {
	r, err := q.db.ExecContext(ctx, createTxn, args.From_account_id, args.To_account_id, args.Amount)
	var i Txn
	if err != nil {
		return i, err
	}
	lastInsertedId, err := r.LastInsertId()
	if err != nil {
		return i, err
	}
	row := q.db.QueryRowContext(ctx, getLastInserted("txns", lastInsertedId))
	err = row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.CreatedAt,
		&i.Amount,
	)
	return i, err
}

const deleteTxn = `-- name: DeleteTxn :exec
DELETE FROM txns
WHERE id = ?
`

func (q *Queries) DeleteTxn(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTxn, id)
	return err
}

const getTxn = `-- name: GetTxn :one
SELECT id, from_account_id, to_account_id, created_at, amount FROM txns
WHERE id = ?
LIMIT 1
`

func (q *Queries) GetTxn(ctx context.Context, id int64) (Txn, error) {
	row := q.db.QueryRowContext(ctx, getTxn, id)
	var i Txn
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.CreatedAt,
		&i.Amount,
	)
	return i, err
}

const getTxns = `-- name: GetTxns :many
SELECT id, from_account_id, to_account_id, created_at, amount FROM txns LIMIT ?, ?
`
type GetTxnArgs struct {
	limit int
	offset int
}
func (q *Queries) GetTxns(ctx context.Context, args GetTxnArgs) ([]Txn, error) {
	rows, err := q.db.QueryContext(ctx, getTxns, args.limit, args.offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Txn
	for rows.Next() {
		var i Txn
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.CreatedAt,
			&i.Amount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTxnAmount = `-- name: UpdateTxnAmount :one
UPDATE txns SET amount = ? WHERE id = ?
`

type UpdateTxnArgs struct {
	Amount int64 `json:"amount" binding:"required"`
	Id int64 `json:"id" binding:"required"`
}

func (q *Queries) UpdateTxnAmount(ctx context.Context, args UpdateTxnArgs) (Txn, error) {
	_, err := q.db.ExecContext(ctx, updateTxnAmount, args.Amount, args.Id)
	var i Txn
	if err != nil {
		return i, err
	}
	row := q.db.QueryRowContext(ctx, getLastInserted("txns", args.Id))
	err = row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.CreatedAt,
		&i.Amount,
	)
	return i, err
}
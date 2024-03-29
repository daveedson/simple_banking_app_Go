// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transfer.sql

package db

import (
	"context"
	"time"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
from_account_id,
to_account_id,
amount,
created_at
) VALUES (
  $1, $2, $3, $4
)RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
		arg.CreatedAt,
	)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, owner, balance, currency, created_at FROM accounts
WHERE id = $1 
LIMIT 1
OFFSET $2
`

type GetTransferParams struct {
	ID     int64 `json:"id"`
	Offset int64 `json:"offset"`
}

func (q *Queries) GetTransfer(ctx context.Context, arg GetTransferParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, arg.ID, arg.Offset)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfer = `-- name: ListTransfer :many
SELECT id, owner, balance, currency, created_at FROM accounts
ORDER BY id
`

func (q *Queries) ListTransfer(ctx context.Context) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listTransfer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
			&i.CreatedAt,
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

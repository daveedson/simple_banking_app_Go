// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: transfer.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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
	FromAccountID int64            `json:"from_account_id"`
	ToAccountID   int64            `json:"to_account_id"`
	Amount        int64            `json:"amount"`
	CreatedAt     pgtype.Timestamp `json:"created_at"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRow(ctx, createTransfer,
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
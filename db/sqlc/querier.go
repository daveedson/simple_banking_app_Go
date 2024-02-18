// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
)

type Querier interface {
	CreatAccount(ctx context.Context, arg CreatAccountParams) (Account, error)
	CreatEntry(ctx context.Context, arg CreatEntryParams) (Account, error)
	CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetEntry(ctx context.Context, id int64) (Entry, error)
	GetTransfer(ctx context.Context, arg GetTransferParams) (Account, error)
	ListAccount(ctx context.Context, arg ListAccountParams) ([]Account, error)
	ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
	ListTransfer(ctx context.Context) ([]Account, error)
	UpdateAuthorBios(ctx context.Context, arg UpdateAuthorBiosParams) (Account, error)
}

var _ Querier = (*Queries)(nil)

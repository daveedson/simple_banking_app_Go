package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {

	arguments := CreatAccountParams{
		Owner:    "Thomas",
		Balance:  100,
		Currency: "USD",
	}

	account, err := testQueries.CreatAccount(context.Background(), arguments)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arguments.Owner, account.Owner)
	require.Equal(t, arguments.Balance, account.Balance)
	require.Equal(t, arguments.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
}

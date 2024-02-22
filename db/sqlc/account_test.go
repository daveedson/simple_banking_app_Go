package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// helper function for creating an account in the test.
func CreateRandomAccount(t *testing.T) Account {

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

	return account
}
func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	firstAccount := CreateRandomAccount(t)
	acc, err := testQueries.GetAccount(context.Background(), firstAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, acc.Balance, firstAccount.Balance)
	require.Equal(t, acc.ID, firstAccount.ID)
	require.Equal(t, acc.Currency, firstAccount.Currency)
	require.WithinDuration(t, acc.CreatedAt, firstAccount.CreatedAt, time.Second)
	require.Equal(t, acc.Owner, firstAccount.Owner)

}

func TestUpdateAccount(t *testing.T) {
	account := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		Balance: account.Balance,
		ID:      account.ID,
	}

	acc, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, acc)

	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, acc.ID, account.ID)
	require.Equal(t, acc.Currency, account.Currency)
	require.WithinDuration(t, acc.CreatedAt, account.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t,account2)

}

func TestListAccount(t *require.TestingT){
	
}

package db

import (
	"context"
	"database/sql"
	"time"

	"testing"

	"github.com/AbdulkarimOgaji/kkmoney3/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	args := CreateAccountArgs{
		Holder:   util.RandomOwner(),
		Currency: util.RandomCurrency(),
	}
	account, err := testingQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, args.Holder, account.Holder)
	require.Equal(t, args.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	args := GetAccountsArgs{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testingQueries.GetAccounts(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, args.Limit)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}

func TestGetAccount(t *testing.T) {
	randomAccount := createRandomAccount(t)
	account, err := testingQueries.GetAccount(context.Background(), randomAccount.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, randomAccount.Balance, account.Balance)
	require.Equal(t, randomAccount.Holder, account.Holder)
	require.Equal(t, randomAccount.ID, account.ID)
	require.Equal(t, randomAccount.Currency, account.Currency)
	require.WithinDuration(t, randomAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	args := UpdateAccountArgs{
		Balance:  100,
		Currency: util.RandomCurrency(),
		Id:       account.ID,
	}
	updatedAccount, err := testingQueries.UpdateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)
	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Holder, updatedAccount.Holder)
	require.Equal(t, args.Balance, updatedAccount.Balance)
	require.Equal(t, args.Currency, updatedAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, updatedAccount.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testingQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	deletedAccount, err := testingQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, deletedAccount)
}

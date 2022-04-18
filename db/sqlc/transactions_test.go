package db

import (
	"context"
	"database/sql"
	"time"

	"testing"

	"github.com/AbdulkarimOgaji/kkmoney3/util"
	"github.com/stretchr/testify/require"
)

func createRandomTxn(t *testing.T) Txn {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	args := CreateTxnArgs{
		From_account_id: account1.ID,
		To_account_id:   account2.ID,
		Amount:          util.RandomBalance(),
	}
	txn, err := testingQueries.CreateTxn(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, txn)

	require.Equal(t, account1.ID, txn.FromAccountID)
	require.Equal(t, account2.ID, txn.ToAccountID)
	require.Equal(t, args.Amount, txn.Amount)
	require.NotZero(t, txn.ID)
	require.NotZero(t, txn.CreatedAt)
	return txn
}

func TestCreateTxn(t *testing.T) {
	createRandomTxn(t)
}

func TestGetTxns(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTxn(t)
	}
	args := GetTxnArgs{
		limit:  5,
		offset: 5,
	}
	txns, err := testingQueries.GetTxns(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, txns)
	require.Len(t, txns, args.limit)

	for _, account := range txns {
		require.NotEmpty(t, account)
	}

}

func TestGetTxn(t *testing.T) {
	randomTxn := createRandomTxn(t)
	txn, err := testingQueries.GetTxn(context.Background(), randomTxn.ID)
	require.NoError(t, err)
	require.NotEmpty(t, txn)
	require.Equal(t, randomTxn.Amount, txn.Amount)
	require.Equal(t, randomTxn.FromAccountID, txn.FromAccountID)
	require.Equal(t, randomTxn.ToAccountID, txn.ToAccountID)
	require.Equal(t, randomTxn.ID, txn.ID)
	require.WithinDuration(t, randomTxn.CreatedAt, txn.CreatedAt, time.Second)
}

func TestUpdateTxn(t *testing.T) {
	txn := createRandomTxn(t)

	args := UpdateTxnArgs{
		Amount: util.RandomBalance(),
		Id:     txn.ID,
	}
	updatedTxn, err := testingQueries.UpdateTxnAmount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTxn)
	require.Equal(t, txn.ID, updatedTxn.ID)
	require.Equal(t, txn.FromAccountID, updatedTxn.FromAccountID)
	require.Equal(t, txn.ToAccountID, updatedTxn.ToAccountID)
	require.Equal(t, args.Amount, updatedTxn.Amount)
	require.WithinDuration(t, txn.CreatedAt, updatedTxn.CreatedAt, time.Second)

}

func TestDeleteTxn(t *testing.T) {
	txn := createRandomTxn(t)
	err := testingQueries.DeleteTxn(context.Background(), txn.ID)
	require.NoError(t, err)

	deletedTxn, err := testingQueries.GetTxn(context.Background(), txn.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, deletedTxn)
}

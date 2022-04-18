package db

import (
	"context"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTxn(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	amount := int64(100)
	n := 1
	results := make(chan TransferResult, n)
	errs := make(chan error, n)

	args := TransferArgs{
		Amount:        amount,
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
	}
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			result, err := store.TransferTx(context.Background(), args)
			results <- result
			errs <- err
		}()
	}
	wg.Wait()
	close(results)
	close(errs)

	for i := 0; i < n; i++ {
		result := <-results
		err := <-errs
		require.NoError(t, err)
		require.NotEmpty(t, result)
		transfer := result.Txn
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)

		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTxn(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := result.FromEntry

		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotEmpty(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry

		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotEmpty(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)
		// check if the accounts have been updated
		require.NotEmpty(t, result.FromAccount)
		require.NotEmpty(t, result.ToAccount)
		diff1 := account1.Balance - result.FromAccount.Balance
		require.Equal(t, diff1, amount)
		diff2 := result.ToAccount.Balance - account2.Balance

		require.Equal(t, diff2, amount)
		require.Equal(t, diff1, diff2)
	}

}

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/AbdulkarimOgaji/kkmoney3/util"

	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account := createRandomAccount(t)
	args := CreateEntryArgs{
		Account_id: account.ID,
		Amount:     util.RandomBalance(),
	}
	entry, err := testingQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)
	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}
	args := GetEntriesArgs{
		Limit:  5,
		Offset: 5,
	}
	entries, err := testingQueries.GetEntries(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entries)
	require.Len(t, entries, args.Limit)

	for _, account := range entries {
		require.NotEmpty(t, account)
	}

}

func TestGetEntry(t *testing.T) {
	randomEntry := createRandomEntry(t)
	entry, err := testingQueries.GetEntry(context.Background(), randomEntry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, randomEntry.Amount, entry.Amount)
	require.Equal(t, randomEntry.AccountID, entry.AccountID)
	require.Equal(t, randomEntry.ID, entry.ID)
	require.WithinDuration(t, randomEntry.CreatedAt, entry.CreatedAt, time.Second)
}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)

	args := UpdateEntryArgs{
		Amount: util.RandomBalance(),
		Id:     entry.ID,
	}
	updatedEntry, err := testingQueries.UpdateEntryAmount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)
	require.Equal(t, entry.ID, updatedEntry.ID)
	require.Equal(t, entry.AccountID, updatedEntry.AccountID)
	require.Equal(t, args.Amount, updatedEntry.Amount)
	require.WithinDuration(t, entry.CreatedAt, updatedEntry.CreatedAt, time.Second)

}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := testingQueries.DeleteEntry(context.Background(), entry.ID)
	require.NoError(t, err)

	deletedEntry, err := testingQueries.GetEntry(context.Background(), entry.ID)
	require.Error(t, err)
	require.ErrorIs(t, err, sql.ErrNoRows)
	require.Empty(t, deletedEntry)
}

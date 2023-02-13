package db

import (
	"context"
	"testing"
	"time"

	"github.com/Bill-Yuyi/Simple-Bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T, account Account) Entry {
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	transfer, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.AccountID, transfer.AccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateEntries(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntries(t, account)
}

func TestGetEntries(t *testing.T) {
	account := createRandomAccount(t)
	entry := createRandomEntries(t, account)
	entry1, err := testQueries.GetEntries(context.Background(), entry.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, entry.ID, entry1.ID)
	require.Equal(t, entry.AccountID, entry1.AccountID)
	require.Equal(t, entry.Amount, entry1.Amount)
	require.WithinDuration(t, entry.CreatedAt, entry.CreatedAt, time.Second)

}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntries(t, account)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}

}

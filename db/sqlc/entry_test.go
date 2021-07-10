package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/puppy002/simple_bank/util"

	"github.com/stretchr/testify/require"
)

// creatRadomentry
func creatRadomEntry(t *testing.T) Entry {
	//创建一批随机账号
	for i := 0; i < 10; i++ {
		creatRadomAccount(t)
	}
	//获取随机账号
	arg1 := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg1)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	arg2 := CreateEntryParams{
		AccountID: accounts[util.RandInt(0, 4)].ID,
		Amount:    util.RandMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, arg2.AccountID, entry.AccountID)
	require.Equal(t, arg2.Amount, entry.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)
	return entry
}

func TestCreateEntry(t *testing.T) {
	creatRadomEntry(t)
}

//test Getentry
func TestGetEntry(t *testing.T) {
	entry1 := creatRadomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)

	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)

}

//test Updateentry
func TestUpdateEntry(t *testing.T) {
	entry1 := creatRadomEntry(t)
	arg := UpdateEntryParams{
		ID:     entry1.ID,
		Amount: util.RandMoney(),
	}
	entry2, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, arg.Amount, entry2.Amount)

	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

//test Delete

func TestDeleteEntry(t *testing.T) {
	entry1 := creatRadomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry1.ID)
	require.NoError(t, err)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, entry2)
}

//test Listentries

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		creatRadomEntry(t)
	}
	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

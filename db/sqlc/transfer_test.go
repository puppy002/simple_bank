package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/puppy002/simple_bank/util"

	"github.com/stretchr/testify/require"
)

// creatRadomAccount
func creatRadomTransfer(t *testing.T) Transfer {
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

	arg2 := CreateTransferParams{
		FromAccountID: accounts[util.RandInt(0, 4)].ID,
		ToAccountID:   accounts[util.RandInt(0, 4)].ID,
		Amount:        util.RandMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg2.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg2.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg2.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfer(t *testing.T) {
	creatRadomTransfer(t)
}

//test GetAccount
func TestGetTransfer(t *testing.T) {
	transfer1 := creatRadomTransfer(t)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)

}

//test UpdateTransfer
func TestUpdateTransfer(t *testing.T) {
	transfer1 := creatRadomTransfer(t)
	arg := UpdateTransferParams{
		ID:     transfer1.ID,
		Amount: util.RandMoney(),
	}
	transfer2, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)

	require.Equal(t, arg.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

//test Delete

func TestDeleteTransfer(t *testing.T) {
	transfer1 := creatRadomTransfer(t)
	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, transfer2)
}

//test ListTransfers

func TestListTransfers(t *testing.T) {
	for i := 0; i < 10; i++ {
		creatRadomTransfer(t)
	}
	arg := ListTransfersParams{
		Limit:  5,
		Offset: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}

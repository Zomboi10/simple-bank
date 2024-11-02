package sqlc

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Zomboi10/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.FromAccountID)
	require.NotZero(t, transfer.ToAccountID)
	require.NotZero(t, transfer.Amount)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	createdTransfer := createRandomTransfer(t)

	transfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, createdTransfer.ID, transfer.ID)
	require.Equal(t, createdTransfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, createdTransfer.Amount, transfer.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt, transfer.CreatedAt, time.Second)
}

func TestUpdateTransFer(t *testing.T) {
	createdTransfer := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:            createdTransfer.ID,
		FromAccountID: createdTransfer.FromAccountID,
		ToAccountID:   createdTransfer.ToAccountID,
		Amount:        util.RandomMoney(),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedTransfer)

	require.Equal(t, createdTransfer.ID, updatedTransfer.ID)
	require.Equal(t, createdTransfer.FromAccountID, updatedTransfer.FromAccountID)
	require.Equal(t, createdTransfer.ToAccountID, updatedTransfer.ToAccountID)
	require.Equal(t, arg.Amount, updatedTransfer.Amount)
	require.WithinDuration(t, createdTransfer.CreatedAt, updatedTransfer.CreatedAt, time.Second)
}

func TestDeleteTransfer(t *testing.T) {
	createdTransfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), createdTransfer.ID)
	require.NoError(t, err)

	getDeletedTransfer, err := testQueries.GetTransfer(context.Background(), createdTransfer.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, getDeletedTransfer)
}

func TestListTransfer(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransfer(t)
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

package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n concurrent transfer transactions
	n := 5
	amount := sql.NullInt64{
		Int64: 10,
		Valid: true,
	}

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.AccountID,
				ToAccountID:   account2.AccountID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		result := <-results

		require.NoError(t, err)
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.AccountID, transfer.FromAccountID)
		require.Equal(t, account2.AccountID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.TransferID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.TransferID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.AccountID, fromEntry.AccountID)
		min_amount := sql.NullInt64{
			Int64: -amount.Int64,
			Valid: true,
		}
		require.Equal(t, min_amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.EntryID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.EntryID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.AccountID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.EntryID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.EntryID)
		require.NoError(t, err)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.AccountID, fromAccount.AccountID)
		require.Equal(t, account1.Currency, fromAccount.Currency)
		require.Equal(t, account1.Owner, fromAccount.Owner)
		require.NotZero(t, fromAccount.CreatedAt)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.AccountID, toAccount.AccountID)
		require.Equal(t, account2.Currency, toAccount.Currency)
		require.Equal(t, account2.Owner, toAccount.Owner)
		require.NotZero(t, toAccount.CreatedAt)

		// check account's balance
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount.Int64 == 0)

		k := int(diff1 / amount.Int64)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balance
	updateAccount1, err := testQueries.GetAccount(context.Background(), account1.AccountID)
	require.NoError(t, err)

	updateAccount2, err := testQueries.GetAccount(context.Background(), account2.AccountID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance-int64(n)*amount.Int64, updateAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount.Int64, updateAccount2.Balance)
}

func TestTransferTxDeadlock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 10
	amount := sql.NullInt64{
		Int64: 10,
		Valid: true,
	}
	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.AccountID
		toAccountID := account2.AccountID

		if i%2 == 1 {
			fromAccountID = account2.AccountID
			toAccountID = account1.AccountID
		}

		go func() {
			_, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check the final updated balance
	updatedAccount1, err := store.GetAccount(context.Background(), account1.AccountID)
	require.NoError(t, err)

	updatedAccount2, err := store.GetAccount(context.Background(), account2.AccountID)
	require.NoError(t, err)

	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}

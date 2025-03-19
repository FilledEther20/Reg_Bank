package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/FilledEther20/Reg_Bank/util"
	"github.com/stretchr/testify/require"
)

// Has been separated so that the running of 1 unit test doesnot effect the result of another also this function would be used in mostly all of the unit tests.
func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(6),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)

	receivedAccount, err := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, err) // To ensure no error returned

	// The below assertion might fail if the timestamp creates issue
	// require.Equal(t, account, receivedAccount)

	require.Equal(t, account.ID, receivedAccount.ID)
	require.Equal(t, account.Owner, receivedAccount.Owner)
	require.Equal(t, account.Balance, receivedAccount.Balance)
	require.Equal(t, account.Currency, receivedAccount.Currency)
	require.WithinDuration(t, account.CreatedAt, receivedAccount.CreatedAt, time.Second) //To ensure the minor differences are considered while checking
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)
	newBalance := util.RandomBalance()
	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: newBalance,
	}

	err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	updatedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.Equal(t, updatedAccount.Balance, newBalance)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account.ID)

	// This may pass even when there is a silent query failure in DB so we have to also verify it.
	require.NoError(t, err)

	// Verifying the operation passed
	deletedAccount, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.Empty(t, deletedAccount)
}

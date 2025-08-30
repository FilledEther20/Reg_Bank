package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/FilledEther20/Reg_Bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		Username:          util.RandomOwner(6),
		HashedPassword:    util.RandomString(8),
		FullName:          util.RandomOwner(12),
		Email:             util.RandomEmail(9),
		PasswordChangedAt: time.Now(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Email, user.Email)

	require.WithinDuration(t, arg.PasswordChangedAt, user.PasswordChangedAt, time.Second)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user := createRandomUser(t)

	receivedUser, err := testQueries.GetUser(context.Background(), user.Username)

	require.NoError(t, err)
	require.NotEmpty(t, receivedUser)

	require.Equal(t, receivedUser.Username, user.Username)
	require.Equal(t, receivedUser.HashedPassword, user.HashedPassword)
	require.Equal(t, receivedUser.Email, user.Email)
	require.Equal(t, receivedUser.FullName, user.FullName)
	require.WithinDuration(t, receivedUser.PasswordChangedAt, user.PasswordChangedAt, time.Second)

}



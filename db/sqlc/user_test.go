package sqlc

import (
	"context"
	"testing"
	"time"

	"github.com/FilledEther20/Reg_Bank/util"
	"github.com/stretchr/testify/require"
)

// Helper to create random user
func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		Username:          util.RandomOwner(6),
		HashedPassword:    hashedPassword,
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

	require.Equal(t, user.Username, receivedUser.Username)
	require.Equal(t, user.HashedPassword, receivedUser.HashedPassword)
	require.Equal(t, user.Email, receivedUser.Email)
	require.Equal(t, user.FullName, receivedUser.FullName)
	require.WithinDuration(t, user.PasswordChangedAt, receivedUser.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user.CreatedAt, receivedUser.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user := createRandomUser(t)

	newFullName := util.RandomOwner(10)
	newEmail := util.RandomEmail(6)

	arg := UpdateUserParams{
		Username: user.Username,
		FullName: newFullName,
		Email:    newEmail,
	}

	updatedUser, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedUser)

	require.Equal(t, user.Username, updatedUser.Username) // username should stay same
	require.Equal(t, newFullName, updatedUser.FullName)
	require.Equal(t, newEmail, updatedUser.Email)
	require.Equal(t, user.HashedPassword, updatedUser.HashedPassword) // password unchanged
}

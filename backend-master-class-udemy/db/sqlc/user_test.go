package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"tutorial.sqlc.dev/app/utils"

	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := utils.HashPassword(utils.RandomPassword())

	arg := CreateUserParams{
		Username:       utils.RandomOwner(),
		FullName:       utils.RandomString(10),
		Email:          utils.RandomEmail(),
		HashedPassword: hashedPassword,
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)

	require.WithinDuration(t, user.CreatedAt, time.Now(), time.Second)
	require.Equal(t, user.PasswordChangedAt, sql.NullTime{
		Time:  time.Time{},
		Valid: false,
	})

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.PasswordChangedAt, user2.PasswordChangedAt)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user1 := createRandomUser(t)

	arg := UpdateUserParams{
		Username:       user1.Username,
		FullName:       utils.RandomString(10),
		Email:          utils.RandomEmail(),
		HashedPassword: utils.RandomPassword(),
		PasswordChangedAt: sql.NullTime{
			Time:  time.Now().In(time.UTC),
			Valid: true,
		},
	}

	user2, err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.FullName, user2.FullName)
	require.Equal(t, arg.Email, user2.Email)
	require.Equal(t, arg.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user1.CreatedAt.UTC(), user2.CreatedAt.UTC(), time.Second)
	require.WithinDuration(t, arg.PasswordChangedAt.Time.In(time.Local), user2.PasswordChangedAt.Time.In(time.Local), time.Second)
}

func TestDeleteUser(t *testing.T) {
	user1 := createRandomUser(t)

	err := testQueries.DeleteUser(context.Background(), user1.Username)
	require.NoError(t, err)

	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.Error(t, err)
	require.Empty(t, user2)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestListUsers(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomUser(t)
	}

	arg := ListUsersParams{
		Limit:  5,
		Offset: 5,
	}

	users, err := testQueries.ListUsers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, users, 5)

	for _, user := range users {
		require.NotEmpty(t, user)
	}
}

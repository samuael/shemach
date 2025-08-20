package user

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/samuael/shemach/backend/pkg/storage/pgx_storage"
	"github.com/samuael/shemach/backend/pkg/storage/pgxconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subosito/gotenv"
)

var userService *UserService

func TestMain(m *testing.M) {
	err := gotenv.Load("../../cmd/main/.env")
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	db, err := pgxconn.NewStorage(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	userRepo := &pgx_storage.UserRepo{
		DB: db,
	}
	userService = (NewUserService(userRepo)).(*UserService)
	os.Exit(m.Run())
}

func TestGetUserByPhone(t *testing.T) {
	t.Parallel()
	result, err := userService.GetUserByPhone(context.Background(), "251911837412")
	require.NoError(t, err)
	assert.NotNil(t, result)
}

func TestDeleteExpiredForgotPasswordSession(t *testing.T) {
	t.Parallel()
	ts := time.Now().Add(-time.Hour)
	_, err := userService.DeleteExpiredForgotPasswordSession(context.Background(), uint64(ts.UnixMilli()))
	require.NoError(t, err)
}

func TestSaveForgotPasswordDetail(t *testing.T) {
	t.Parallel()
	status, shortCOde, err := userService.SaveForgotPasswordDetail(context.Background(), "251911837412", "ABD12", true, false, false, 10)
	require.NoError(t, err)
	assert.NotZero(t, status)
	println(status)
	assert.Equal(t, "ABD12", shortCOde)
}

func TestConfirmForgotPasswordShortcode(t *testing.T) {
	t.Parallel()
	status, shortCOde, err := userService.SaveForgotPasswordDetail(context.Background(), "251911837412", "ABD12", true, false, false, 10)
	require.NoError(t, err)
	require.True(t, status > 3)

	statusCode, err := userService.ConfirmForgotPasswordShortcode(context.Background(), "251911837412", shortCOde)
	require.NoError(t, err)
	assert.Equal(t, int64(0), statusCode)
}

func TestChangeUserPassword(t *testing.T) {
	t.Parallel()
	err := userService.ChangeUserPassword(context.Background(), 1743239668204, "sfdajhfakdjfakjdfhaskd")
	require.NoError(t, err)
}

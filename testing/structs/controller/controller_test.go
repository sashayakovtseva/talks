package controller

import (
	"context"
	"testing"

	"fmt"

	"github.com/sashayakovtseva/talks/testing/structs/database"
	"github.com/stretchr/testify/require"
)

type (
	errorUserProvider      struct{}
	notManagerUserProvider struct{}
	okUserProvider         struct{}
)

func (errorUserProvider) Fetch(ctx context.Context, id int64) (database.User, error) {
	return database.User{}, fmt.Errorf("always error")
}

func (notManagerUserProvider) Fetch(ctx context.Context, id int64) (database.User, error) {
	return database.User{ID: id, IsManager: false}, nil
}

func (okUserProvider) Fetch(ctx context.Context, id int64) (database.User, error) {
	return database.User{ID: id, IsManager: true}, nil
}

func TestFetchUser(t *testing.T) {
	tt := []struct {
		name          string
		usersProvider UsersProvider
		expectResult  *database.User
		expectError   error
	}{
		{
			name:          "fetch user error",
			usersProvider: errorUserProvider{},
			expectResult:  nil,
			expectError:   fmt.Errorf("could not fetch manager: always error"),
		},
		{
			name:          "not manager",
			usersProvider: notManagerUserProvider{},
			expectResult:  nil,
			expectError:   fmt.Errorf("view user not allowed"),
		},
		{
			name:          "all ok",
			usersProvider: okUserProvider{},
			expectResult:  &database.User{ID: 0, IsManager: true},
			expectError:   nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := userController{
				usersProvider: tc.usersProvider,
			}
			actual, err := c.FetchUser(context.Background(), 0, 0)
			require.Equal(t, tc.expectResult, actual)
			require.Equal(t, tc.expectError, err)
		})
	}
}

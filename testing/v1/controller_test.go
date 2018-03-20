package v1

import (
	"context"
	"fmt"
	"testing"

	"github.com/sashayakovtseva/talks/testing/database"
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

func (errorUserProvider) IsManager(ctx context.Context, id int64) (bool, error) {
	return false, fmt.Errorf("always error")
}

func (notManagerUserProvider) Fetch(ctx context.Context, id int64) (database.User, error) {
	return database.User{ID: id}, nil
}

func (notManagerUserProvider) IsManager(ctx context.Context, id int64) (bool, error) {
	return false, nil
}

func (okUserProvider) Fetch(ctx context.Context, id int64) (database.User, error) {
	return database.User{ID: id}, nil
}

func (okUserProvider) IsManager(ctx context.Context, id int64) (bool, error) {
	return true, nil
}

func TestFetchUser(t *testing.T) {
	tt := []struct {
		name          string
		usersProvider usersProvider
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
			expectResult:  &database.User{ID: 0},
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

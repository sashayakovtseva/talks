package v2

import (
	"context"
	"fmt"
	"testing"

	"github.com/sashayakovtseva/talks/testing/database"
	"github.com/stretchr/testify/require"
)

type (
	fetchFunc     func(ctx context.Context, id int64) (database.User, error)
	isManagerFunc func(ctx context.Context, id int64) (bool, error)

	testUserProvider struct {
		fetch     fetchFunc
		isManager isManagerFunc
	}
)

func (t testUserProvider) Fetch(ctx context.Context, id int64) (database.User, error) {
	return t.fetch(ctx, id)
}

func (t testUserProvider) IsManager(ctx context.Context, id int64) (bool, error) {
	return t.isManager(ctx, id)
}

func errorFetch(_ context.Context, _ int64) (database.User, error) {
	return database.User{}, fmt.Errorf("always error")
}

func isNotManager(_ context.Context, _ int64) (bool, error) {
	return false, nil
}

func isManager(_ context.Context, _ int64) (bool, error) {
	return true, nil
}

func errorIsManager(_ context.Context, _ int64) (bool, error) {
	return false, fmt.Errorf("always error")
}

func okFetch(_ context.Context, id int64) (database.User, error) {
	return database.User{ID: id}, nil
}

func TestFetchUser(t *testing.T) {
	tt := []struct {
		name          string
		fetchFunc     fetchFunc
		isManagerFunc isManagerFunc
		expectResult  *database.User
		expectError   error
	}{
		{
			name:          "fetch manager error",
			isManagerFunc: errorIsManager,
			expectResult:  nil,
			expectError:   fmt.Errorf("could not fetch manager: always error"),
		},
		{
			name:          "not manager",
			isManagerFunc: isNotManager,
			expectResult:  nil,
			expectError:   fmt.Errorf("view user not allowed"),
		},
		{
			name:          "fetch user error",
			isManagerFunc: isManager,
			fetchFunc:     errorFetch,
			expectResult:  nil,
			expectError:   fmt.Errorf("could not fetch user: always error"),
		},
		{
			name:          "all ok",
			isManagerFunc: isManager,
			fetchFunc:     okFetch,
			expectResult:  &database.User{ID: 0},
			expectError:   nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := userController{
				usersProvider: testUserProvider{
					fetch:     tc.fetchFunc,
					isManager: tc.isManagerFunc,
				},
			}
			actual, err := c.FetchUser(context.Background(), 0, 0)
			require.Equal(t, tc.expectResult, actual)
			require.Equal(t, tc.expectError, err)
		})
	}
}

package v2

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/sashayakovtseva/talks/testing/database"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type (
	testUserProvider struct {
		mock.Mock
	}
)

var errMock = fmt.Errorf("")

func (t *testUserProvider) Fetch(ctx context.Context, id int64) (database.User, error) {
	args := t.Called(ctx, id)
	return args.Get(0).(database.User), args.Error(1)
}

func (t *testUserProvider) IsManager(ctx context.Context, id int64) (bool, error) {
	args := t.Called(ctx, id)
	return args.Get(0).(bool), args.Error(1)
}

func TestFetchUser(t *testing.T) {
	ctx := context.Background()

	userProviderMock := new(testUserProvider)
	defer userProviderMock.AssertExpectations(t)

	userProviderMock.On("IsManager", ctx, int64(0)).Return(false, errMock)
	userProviderMock.On("IsManager", ctx, int64(1)).Return(false, nil)
	userProviderMock.On("IsManager", ctx, mock.Anything).Return(true, nil)
	userProviderMock.On("ByID", ctx, int64(0)).Return((*database.User)(nil), errMock)
	userProviderMock.On("ByID", ctx, int64(1)).Return(&database.User{ID: 1, Name: "sasha"}, errMock)

	tt := []struct {
		name         string
		managerID    int64
		userID       int64
		expectResult *database.User
		expectError  error
	}{
		{
			name:        "fetch manager error",
			managerID:   0,
			expectError: fmt.Errorf("could not fetch manager: %v", errMock),
		},
		{
			name:        "not manager",
			managerID:   1,
			expectError: fmt.Errorf("view user not allowed"),
		},
		{
			name:        "fetch user error",
			managerID:   2,
			userID:      0,
			expectError: fmt.Errorf("could not fetch user: %v", errMock),
		},
		{
			name:         "all ok",
			managerID:    2,
			userID:       1,
			expectError:  nil,
			expectResult: &database.User{ID: 1, Name: "sasha"},
		},
	}

	c := userController{
		usersProvider: userProviderMock,
	}
	var wg sync.WaitGroup
	wg.Add(len(tt))
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			defer wg.Done()
			actual, err := c.FetchUser(ctx, tc.managerID, tc.userID)
			require.Equal(t, tc.expectResult, actual)
			require.Equal(t, tc.expectError, err)
		})
	}
}

package controller

import (
	"context"
	"fmt"

	"github.com/sashayakovtseva/talks/testing/structs/database"
)

type (
	UsersProvider interface {
		Fetch(ctx context.Context, id int64) (database.User, error)
	}

	userController struct {
		usersProvider UsersProvider
	}
)

var (
	Users = userController{
		usersProvider: database.Users,
	}
)

func (u userController) FetchUser(ctx context.Context, managerID, userID int64) (*database.User, error) {
	manager, err := u.usersProvider.Fetch(ctx, managerID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch manager: %v", err)
	}
	if !manager.IsManager {
		return nil, fmt.Errorf("view user not allowed")
	}
	user, err := u.usersProvider.Fetch(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch user: %v", err)
	}
	return &user, nil
}

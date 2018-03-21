package v2

import (
	"context"
	"fmt"

	"github.com/sashayakovtseva/talks/testing/database"
)

type (
	usersProvider interface {
		Fetch(ctx context.Context, id int64) (database.User, error)
		IsManager(ctx context.Context, id int64) (bool, error)
	}

	userController struct {
		usersProvider usersProvider
	}
)

var (
	// Users provide access to users business logic
	Users = userController{
		usersProvider: database.Users,
	}
)

// FetchUser checks that viewer has manager permissions and if he has
// fetches the user by id
func (u userController) FetchUser(ctx context.Context, managerID, userID int64) (*database.User, error) {
	isManager, err := u.usersProvider.IsManager(ctx, managerID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch manager: %v", err)
	}
	if !isManager {
		return nil, fmt.Errorf("view user not allowed")
	}
	user, err := u.usersProvider.Fetch(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch user: %v", err)
	}
	return &user, nil
}

package v0

import (
	"context"
	"fmt"

	"github.com/sashayakovtseva/talks/testing/database"
)

// FetchUser checks that viewer has manager permissions and if he has
// fetches the user by id
func FetchUser(ctx context.Context, managerID, userID int64) (*database.User, error) {
	isManager, err := database.Users.IsManager(ctx, managerID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch manager: %v", err)
	}
	if !isManager {
		return nil, fmt.Errorf("view user not allowed")
	}
	user, err := database.Users.Fetch(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch user: %v", err)
	}
	return &user, nil
}

package controller

import (
	"context"
	"fmt"

	"github.com/sashayakovtseva/talks/testing/bad/database"
)

func FetchUser(ctx context.Context, managerID, userID int64) (*database.User, error) {
	manager, err := database.FetchUser(ctx, managerID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch manager: %v", err)
	}
	if !manager.IsManager {
		return nil, fmt.Errorf("view user not allowed")
	}
	user, err := database.FetchUser(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("could not fetch user: %v", err)
	}
	return &user, nil
}

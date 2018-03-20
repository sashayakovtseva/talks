package v0

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchUser(t *testing.T) {
	user, err := FetchUser(context.Background(), 0, 0)
	assert.Nil(t, user)
	assert.Error(t, err)
}

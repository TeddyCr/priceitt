//go:build unit
// +build unit

package serializer

import (
	"testing"

	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestJsonToString(t *testing.T) {
	user := &entities.User{
		ID:   uuid.New(),
		Name: "John Doe",
	}
	jsonBytes, err := JsonToString(user)
	assert.NoError(t, err)
	assert.IsType(t, []byte{}, jsonBytes)
}

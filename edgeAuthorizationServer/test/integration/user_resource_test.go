package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/TeddyCr/priceitt/models/generated/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var resourcePath = "http://localhost:8000/api/v1/user"

func TestMain(m *testing.M) {
	compose := StartApplication()
	m.Run()
	StopApplication(compose)
}

func TestCreateUser(t *testing.T) {
	body := []byte(`{
		"name": "John Doe",
		"email": "john.d@example.com",
		"password": "*lX1t6r8};k}8VPYEk",
		"confirmPassword": "*lX1t6r8};k}8VPYEk",
		"authType": "basic"
	}`)
	bodyReader := bytes.NewReader(body)

	req, err := http.Post(resourcePath, "application/json", bodyReader)
	require.NoError(t, err)
	defer req.Body.Close()

	resp, err := io.ReadAll(req.Body)
	require.NoError(t, err)

	var user entities.User
	err = json.Unmarshal(resp, &user)
	require.NoError(t, err)
	
	assert.NotNil(t, user.ID)
	assert.Equal(t, "John Doe", user.Name)
	assert.NotEqual(t, 0, user.CreatedAt)
	assert.NotEqual(t, 0, user.UpdatedAt)
	authMechanism := user.AuthenticationMechanism.(map[string]interface{})
	assert.Equal(t, "", authMechanism["password"])
}


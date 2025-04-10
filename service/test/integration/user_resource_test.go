//go:build integration
// +build integration

package integration

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/TeddyCr/priceitt/service/models/generated/entities"
	"github.com/go-chi/httplog/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var resourcePath = "http://localhost:8000/api/v1/user"

func TestMain(m *testing.M) {
	compose := StartApplication()
	m.Run()
	StopApplication(compose)
}

func createUser() (*http.Response, error) {
	body := []byte(`{
		"name": "John Doe",
		"email": "john.d@example.com",
		"password": "*lX1t6r8};k}8VPYEk",
		"confirmPassword": "*lX1t6r8};k}8VPYEk",
		"authType": "basic"
	}`)
	bodyReader := bytes.NewReader(body)

	req, err := http.Post(resourcePath, "application/json", bodyReader)
	return req, err
}

func TestCreateUser(t *testing.T) {
	req, err := createUser()
	require.NoError(t, err)
	defer func() {
		if err := req.Body.Close(); err != nil {
			t.Errorf("Error closing response body: %v", err)
		}
	}()

	// Get the logger from the response request context
	logger := httplog.LogEntry(req.Request.Context())
	logger.Debug("Creating user", "email", "john.d@example.com")

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

	logger.Debug("User created successfully", "userId", user.ID)
}

func TestLogin(t *testing.T) {
	createUser() //nolint:errcheck
	body := []byte(`{
		"username": "John Doe",
		"password": "*lX1t6r8};k}8VPYEk"
	}`)
	bodyReader := bytes.NewReader(body)

	req, err := http.Post(resourcePath+"/login", "application/json", bodyReader)
	require.NoError(t, err)
	defer func() {
		if err := req.Body.Close(); err != nil {
			t.Errorf("Error closing response body: %v", err)
		}
	}()

	// Get the logger from the response request context
	logger := httplog.LogEntry(req.Request.Context())
	logger.Debug("Attempting login", "username", "John Doe")

	resp, err := io.ReadAll(req.Body)
	require.NoError(t, err)

	var token entities.JWToken
	err = json.Unmarshal(resp, &token)
	require.NoError(t, err)

	assert.NotNil(t, token.Token)
	assert.NotEqual(t, 0, token.CreatedAt)
	assert.NotEqual(t, 0, token.UpdatedAt)

	logger.Debug("Login successful")
}

func TestLogOut(t *testing.T) {
	createUser() //nolint:errcheck
	body := []byte(`{
		"username": "John Doe",
		"password": "*lX1t6r8};k}8VPYEk"
	}`)
	bodyReader := bytes.NewReader(body)

	req, err := http.Post(resourcePath+"/login", "application/json", bodyReader)
	require.NoError(t, err)
	defer func() {
		if err := req.Body.Close(); err != nil {
			t.Errorf("Error closing response body: %v", err)
		}
	}()

	resp, err := io.ReadAll(req.Body)
	require.NoError(t, err)

	var token entities.JWToken
	err = json.Unmarshal(resp, &token)
	require.NoError(t, err)

	tokenString := token.Token
	tokenString = base64.StdEncoding.EncodeToString([]byte(tokenString))

	newReq, err := http.NewRequest("POST", resourcePath+"/logout", nil)
	require.NoError(t, err)
	newReq.Header.Set("Authorization", "Bearer "+tokenString)
	client := http.DefaultClient
	_, err = client.Do(newReq)
	require.NoError(t, err)
}

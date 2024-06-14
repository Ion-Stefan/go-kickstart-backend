package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ion-Stefan/go-kickstart-backend/types"
	"github.com/gorilla/mux"
)

// Create a mock user store
type mockUserStore struct{}

// Implement the methods for the mock user store
func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(user types.User) error {
  return nil
}

func TestUserRoutes(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	// Add a test case for the user entering bad payload
	t.Run("should fail if the user payload is invalid", func(t *testing.T) {
		// Create a payload with an invalid email
		payload := types.RegisterUserPayload{
			FirstName: "Go",
			LastName:  "Kickstart",
			Email:     "gokickstart-invaidemail",
			Password:  "password",
		}
		// Marshal the payload
		marshalled, _ := json.Marshal(payload)

		// Create a new request
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		// Create a new recorder
		rr := httptest.NewRecorder()
		// Create a new router
		router := mux.NewRouter()

		// Register the handleRegister method
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)

		// Serve the request
		router.ServeHTTP(rr, req)

		// Check if the status code is 400
		if rr.Code != http.StatusBadRequest {
			// If it is not, fail the test
			t.Fatalf("expected status code %d, got %d", http.StatusBadRequest, rr.Code)
		}
	})

	// Add a test case for corectly registering the user
	t.Run("should corectly register the user", func(t *testing.T) {
		// Create a payload with a valid email
		payload := types.RegisterUserPayload{
			FirstName: "Go",
			LastName:  "Kickstart",
			Email:     "gokickstart@gmail.com",
			Password:  "password",
		}
		// Marshal the payload
		marshalled, _ := json.Marshal(payload)

		// Create a new request
		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		// Create a new recorder
		rr := httptest.NewRecorder()
		// Create a new router
		router := mux.NewRouter()

		// Register the handleRegister method
		router.HandleFunc("/register", handler.handleRegister).Methods(http.MethodPost)

		// Serve the request
		router.ServeHTTP(rr, req)

		// Check if the status code is 201
		if rr.Code != http.StatusCreated {
			// If it is not, fail the test
			t.Fatalf("expected status code %d, got %d", http.StatusOK, rr.Code)
		}
	})
}


package user

import (
	"fmt"
	"net/http"

	"github.com/Ion-Stefan/go-kickstart-backend/service/auth"
	"github.com/Ion-Stefan/go-kickstart-backend/types"
	"github.com/Ion-Stefan/go-kickstart-backend/utils"
	"github.com/go-playground/validator/v10"

	"github.com/gorilla/mux"
)

// Struct that holds the methods for handling user requests
type Handler struct {
	store types.UserStore
}

// Creates a new Handler
func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// Register the routes for the user handler
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Create a subrouter for the user routes
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// handleLogin handles the login request
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
}

// handleRegister handles the register request
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get the payload from the request
	var payload types.RegisterUserPayload
	// Parse the JSON payload
	if err := utils.ParseJson(r, &payload); err != nil {
		// If there is an error, write the error to the response
		utils.WriteError(w, http.StatusBadRequest, err)
	}
	// Validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", error))
		return
	}

	// Check if the user already exists
	_, err := h.store.GetUserByEmail(payload.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email: %s already exists", payload.Email))
		return
	}

	hashedPassword, _ := auth.HashPassword(payload.Password)

	// Create a new user
	err = h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

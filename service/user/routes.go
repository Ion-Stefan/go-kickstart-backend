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
	// Get the payload from the request
	var user types.LoginUserPayload
	// Parse the JSON payload
	if err := utils.ParseJson(r, &user); err != nil {
		// If there is an error, write the error to the response
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// Validate the payload
	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid or missing body: %v", errors))
		return
	}

	// Check if the user already exists
	u, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		fmt.Println(u.CreatedAt, u.FirstName, u.LastName, u.Email, u.Password, u.ID)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid"))
		return
	}
	if !auth.ComparePasswords([]byte(u.Password), user.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found or invalid email or invalid password"))
		return
	}
	println(u.Password)
	println(user.Password)

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": ""})
}

// handleRegister handles the register request
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// Get the payload from the request
	var user types.RegisterUserPayload
	// Parse the JSON payload
	if err := utils.ParseJson(r, &user); err != nil {
		// If there is an error, write the error to the response
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// Validate the payload
	if err := utils.Validate.Struct(user); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid or missing body: %v", error))
		return
	}

	// Check if the user already exists
	_, err := h.store.GetUserByEmail(user.Email)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email: %s already exists", user.Email))
		return
	}

	// Hash the password
	hashedPassword, _ := auth.HashPassword(user.Password)

	// Create a new user
	err = h.store.CreateUser(types.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}

package item

import (
	"fmt"
	"net/http"

	"github.com/Ion-Stefan/go-kickstart-backend/service/auth"
	"github.com/Ion-Stefan/go-kickstart-backend/types"
	"github.com/Ion-Stefan/go-kickstart-backend/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store     types.ItemStore
	userStore types.UserStore
}

func NewHandler(store types.ItemStore, userStore types.UserStore) Handler {
	return Handler{store: store, userStore: userStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/items", h.GetItems).Methods("GET")
	// Add the JWT middleware to the route
	router.HandleFunc("/items", auth.WithJWTAuth(h.CreateItem, h.userStore)).Methods("POST")
}

func (h *Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	// Get the items from the store
	items, err := h.store.GetItems()
	if err != nil {
		http.Error(w, "Failed to get items", http.StatusInternalServerError)
		return
	}
	utils.WriteJSON(w, http.StatusOK, items)
}

func (h *Handler) CreateItem(w http.ResponseWriter, r *http.Request) {
	// Get the payload from the request
	var item types.ItemPayload
	// Parse the JSON payload
	if err := utils.ParseJson(r, &item); err != nil {
		// If there is an error, write the error to the response
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// Validate the payload
	if err := utils.Validate.Struct(item); err != nil {
		error := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid or missing body: %v", error))
		return
	}

	// Create a new user
	err := h.store.CreateItem(types.Item{
		Name:        item.Name,
		Description: item.Description,
		ImageURL:    item.ImageURL,
	})
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, fmt.Sprintf("Item: '%s' created", item.Name))
}

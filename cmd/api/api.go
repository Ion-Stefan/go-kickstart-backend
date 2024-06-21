package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Ion-Stefan/go-kickstart-backend/service/item"
	"github.com/Ion-Stefan/go-kickstart-backend/service/user"
	"github.com/gorilla/mux"
)

// Defining the APIServer
type APIServer struct {
	db   *sql.DB
	addr string
}

// Creates the APIServer
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the APIServer
func (s *APIServer) Run() error {
	// Create the router
	router := mux.NewRouter()
	// Create the subrouter for the API
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create the user store
	userStore := user.NewStore(s.db)
	// Create the user handler
	userHandler := user.NewHandler(userStore)
	// Register the user routes
	userHandler.RegisterRoutes(subrouter)

	// Create the item store
	itemStore := item.NewStore(s.db)
	// Create the item handler
	itemHandler := item.NewHandler(itemStore)
	// Register the item routes
	itemHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	// Start the server
	return http.ListenAndServe(s.addr, router)
}

package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Ion-Stefan/go-kickstart-backend/service/user"
	"github.com/gorilla/mux"
)

// Defining the APIServer
type APIServer struct {
	db   *sql.DB
	addr string
}

// Creates a new APIServer
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

// Run starts the APIServer
func (s *APIServer) Run() error {
	// Create a new router
	router := mux.NewRouter()
	// Create a subrouter for the API
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create a new user store
	userStore := user.NewStore(s.db)
	// Create a new user handler
	userHandler := user.NewHandler(userStore)
	// Register the user routes
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on", s.addr)

	// Start the server
	return http.ListenAndServe(s.addr, router)
}

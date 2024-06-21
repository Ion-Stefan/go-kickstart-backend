package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Ion-Stefan/go-kickstart-backend/cmd/api"
	"github.com/Ion-Stefan/go-kickstart-backend/config"
	"github.com/Ion-Stefan/go-kickstart-backend/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	// Load the environment variables from the .env file
	db, err := db.NewMySQLStorge(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Print the environment variables
	fmt.Printf("mysql username: %v\nmysql password: %v\nmysql database name: %v\n\n", config.Envs.DBUser, config.Envs.DBPassword, config.Envs.DBName)
	// Initialize the storage
	initStorage(db)

	// Create the API server on the port specified in the .env file
	server := api.NewAPIServer(fmt.Sprintf(":%v", config.Envs.Port), db)
	// Run the server
	if err := server.Run(); err != nil {
		// If an error occurs, log it
		log.Fatal(err)
	}
}

// Initialize the storage
func initStorage(db *sql.DB) {
	// db.Ping checks if the connection to the database is successful
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB: Successfully connected!")
}

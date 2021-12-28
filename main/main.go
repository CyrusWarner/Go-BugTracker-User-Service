package main

import (
	"github.com/Go-BugTracker-User-Service/db_client"
	"github.com/Go-BugTracker-User-Service/handlers"
)

func main() {
	db_client.InitializeDbConnection()

	defer db_client.DBClient.Close()

	handlers.Router()
}

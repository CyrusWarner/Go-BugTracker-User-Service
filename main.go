package main

import (
	"log"
	"net/http"

	"github.com/Go-BugTracker-User-Service/db_client"
	h "github.com/Go-BugTracker-User-Service/handlers"
	"github.com/gorilla/mux"
)

func main() {
	db_client.InitializeDbConnection()

	defer db_client.DBClient.Close()

	Router()

}

func Router() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user/register", h.UserRegisterHandler).Methods("POST")
	r.HandleFunc("/api/user/login", h.UserLoginHandler).Methods("POST")

	// Allows requests coming from any domain with port 4000. No domain currently so this will be used for testing
	log.Fatal(http.ListenAndServe("0.0.0.0:4000", r))

}

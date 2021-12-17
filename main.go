package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Go-BugTracker-User-Service/db_client"
	"github.com/Go-BugTracker-User-Service/models"

	"github.com/gorilla/mux"
)

func main() {
	db_client.InitializeDbConnection()

	defer db_client.DBClient.Close()

	router()
}

func router() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user/register", handleUserRegister).Methods("POST")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func handleUserRegister(w http.ResponseWriter, r *http.Request) {
	u := models.User{}

	decoder := json.NewDecoder(r.Body) // NewDecoder that returns a new decoder that reads from r
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User Object")
		return
	}
}

func respondWithError(w http.ResponseWriter, statusCode int, errMessage string) {
	respondWithJSON(w, statusCode, map[string]string{"error": errMessage})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

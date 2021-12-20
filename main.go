package main

import (
	"encoding/json"
	"log"
	"net/http"

	usercontroller "github.com/Go-BugTracker-User-Service/controllers"
	"github.com/Go-BugTracker-User-Service/db_client"
	models "github.com/Go-BugTracker-User-Service/models"
	"github.com/gorilla/mux"
)

func main() {
	db_client.InitializeDbConnection()

	defer db_client.DBClient.Close()

	router()
}

func router() {
	r := mux.NewRouter()

	r.HandleFunc("/api/user/register", userRegisterHandler).Methods("POST")
	r.HandleFunc("/api/user/login", userLoginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func userRegisterHandler(w http.ResponseWriter, r *http.Request) {
	u := models.UserRegister{}
	var err error

	decoder := json.NewDecoder(r.Body) // NewDecoder that returns a new decoder that reads from r
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User Object")
		return
	}

	if u, err = usercontroller.RegisterUser(db_client.DBClient, u); err != nil {
		switch err.Error() {
		case usercontroller.ErrUserRegistered.Error():
			respondWithError(w, http.StatusBadRequest, usercontroller.ErrUserRegistered.Error())
			return
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	respondWithJSON(w, http.StatusCreated, u)
}

func userLoginHandler(w http.ResponseWriter, r *http.Request) {
	ul := models.UserLogin{}
	var err error

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ul); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User Login Object")
		return
	}

	ut := models.UserToken{}

	if ut, err = usercontroller.LoginUser(db_client.DBClient, ul); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ut)
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

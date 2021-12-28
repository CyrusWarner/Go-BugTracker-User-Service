package handlers

import (
	"encoding/json"
	"net/http"

	uc "github.com/Go-BugTracker-User-Service/controllers"
	"github.com/Go-BugTracker-User-Service/db_client"
	m "github.com/Go-BugTracker-User-Service/models"
)

func UserRegisterHandler(w http.ResponseWriter, r *http.Request) {
	u := m.UserRegister{}
	var err error

	decoder := json.NewDecoder(r.Body) // NewDecoder that returns a new decoder that reads from r
	if err := decoder.Decode(&u); err != nil {
		respondWithError(w, http.StatusBadRequest, "INVALID_USER_REGISTER_OBJECT")
		return
	}

	defer r.Body.Close()

	if u, err = uc.RegisterUser(db_client.DBClient, u); err != nil {
		switch err.Error() {
		case uc.ErrUserRegistered.Error():
			respondWithError(w, http.StatusBadRequest, uc.ErrUserRegistered.Error())
			return
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	respondWithJSON(w, http.StatusCreated, u)
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	ul := m.UserLogin{}
	var err error

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ul); err != nil {
		respondWithError(w, http.StatusBadRequest, "INVALID_USER_LOGIN_OBJECT")
		return
	}

	defer r.Body.Close()

	ut := m.UserToken{}

	if ut, err = uc.LoginUser(db_client.DBClient, ul); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, ut)
}

func respondWithError(w http.ResponseWriter, statusCode int, errMessage string) {
	respondWithJSON(w, statusCode, map[string]string{"message": errMessage})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	res, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(res)
}

package usermodel

import (
	"database/sql"
	"errors"
	"strings"
	"time"

	jwt "github.com/Go-BugTracker-User-Service/jwt"
	models "github.com/Go-BugTracker-User-Service/models"

	"golang.org/x/crypto/bcrypt" // package used to create hashes of passwords and read hashes of passwords
)

var ErrUserRegistered error = errors.New("user already registered")
var ErrUserLogin error = errors.New("invalid login")

func RegisterUser(db *sql.DB, ur models.UserRegister) (models.UserRegister, error) {
	var err error
	var row *sql.Row
	passwordHash, hashErr := hashPassword(ur.Password)
	if hashErr != nil {
		return ur, hashErr
	}

	ur.Email = strings.ToLower(ur.Email)

	row = db.QueryRow("SELECT * FROM Users WHERE Email=@p1",
		ur.Email,
	)

	userLookup := models.UserRegister{}
	err = row.Scan(&userLookup)

	if err != sql.ErrNoRows {
		return ur, ErrUserRegistered
	}

	row = db.QueryRow("INSERT INTO Users (FirstName, LastName, Email, Password, DateJoined) OUTPUT Inserted.UserId, Inserted.FirstName, Inserted.LastName, Inserted.Email, Inserted.DateJoined Values(@p1, @p2, @p3, @p4, @p5)",
		ur.FirstName,
		ur.LastName,
		ur.Email,
		passwordHash,
		time.Now(),
	)

	err = row.Scan(
		&ur.UserId,
		&ur.FirstName,
		&ur.LastName,
		&ur.Email,
		&ur.DateJoined,
	)

	return ur, err
}

func LoginUser(db *sql.DB, ul models.UserLogin) (models.UserToken, error) {
	row := db.QueryRow("SELECT * FROM Users WHERE Email=@p1",
		ul.Email,
	)

	u := models.User{}
	ut := models.UserToken{}

	err := row.Scan(
		&u.UserId,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.EmailConfirmed,
		&u.DateJoined,
	)
	if err != nil {
		return ut, err
	}

	canLoginWithPassword := checkPasswordHash(ul.Password, u.Password) // Checks to see if password is correct.

	if !canLoginWithPassword {
		return ut, ErrUserLogin
	}

	utd := models.UserTokenData{
		UserId:         u.UserId,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Email:          u.Email,
		EmailConfirmed: u.EmailConfirmed,
		DateJoined:     u.DateJoined,
	}

	tokenString, err := jwt.GenerateJWT(utd)
	if err != nil {
		return ut, err
	}

	ut = models.UserToken{UserToken: tokenString}

	return ut, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password string, passwordhash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordhash), []byte(password))
	return err == nil // if err is equal to nil tha means no errors have occured and passwords match
}

package usermodel

import (
	"database/sql"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt" // package used to create hashes of passwords and read hashes of passwords
)

type UserRegister struct {
	UserId         int    `json:"userId"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Password       string `json:"-"` // omit the password field
	EmailConfirmed bool   `json:",omitempty"`
	DateJoined     string `json:",omitempty"`
}

var ErrUserRegistered error = errors.New("user already registered")

func RegisterUser(db *sql.DB, user UserRegister) (UserRegister, error) {
	var err error
	var row *sql.Row
	passwordHash, hashErr := hashPassword(user.Password)
	if hashErr != nil {
		return user, hashErr
	}

	row = db.QueryRow("SELECT * FROM Users WHERE Email=@p1",
		user.Email,
	)

	userLookup := UserRegister{}
	err = row.Scan(&userLookup)

	if err != sql.ErrNoRows {
		return user, ErrUserRegistered
	}

	row = db.QueryRow("INSERT INTO Users (FirstName, LastName, Email, Password, DateJoined) OUTPUT INSERTED.* Values(@p1, @p2, @p3, @p4, @p5)",
		user.FirstName,
		user.LastName,
		user.Email,
		passwordHash,
		time.Now(),
	)

	err = row.Scan(
		&user.UserId,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.EmailConfirmed,
		&user.DateJoined,
	)

	return user, err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password string, passwordhash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordhash), []byte(password))
	return err == nil // if err is equal to nil tha means no errors have occured and passwords match
}

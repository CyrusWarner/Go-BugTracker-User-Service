package models

import "github.com/golang-jwt/jwt"

type UserRegister struct {
	UserId     int    `json:"userId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Password   string `json:"password"` // omit the password field
	DateJoined string `json:",omitempty"`
}

type UserRegisterData struct {
	UserId     int    `json:"userId"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	DateJoined string `json:",omitempty"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	UserToken string `json:"token"`
}

type UserTokenData struct {
	UserId         int
	FirstName      string
	LastName       string
	Email          string
	EmailConfirmed bool
	DateJoined     string
}

type Claims struct {
	UserTokenData
	jwt.StandardClaims
}
type User struct {
	UserId         int    `json:"UserId"`
	FirstName      string `json:"FirstName"`
	LastName       string `json:"LastName"`
	Email          string `json:"email"`
	Password       string `json:"-"`
	EmailConfirmed bool   `json:"EmailConfirmed"`
	DateJoined     string `json:"DateJoined"`
}

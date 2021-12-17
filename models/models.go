package models

type User struct {
	UserId         string `json:"userId"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	EmailConfirmed bool   `json:"emailConfirmed"`
	DateJoined     string `json:"dateJoined"`
}

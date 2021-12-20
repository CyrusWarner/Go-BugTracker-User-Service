package token

import (
	"log"
	"os"
	"time"

	"github.com/Go-BugTracker-User-Service/models"
	jwt "github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func GenerateJWT(tokenData models.UserTokenData) (models.UserToken, error) {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalln("could not load .env file")
	}
	mySecretKey := os.Getenv("JWTKEY")

	ut := models.UserToken{}

	// TODO CHANGE THIS TO A WEEK LONG
	expirationTime := time.Now().Add(time.Minute * 5).Unix() // create the expiration time for the token
	claims := &models.Claims{                                // creates the claims struct to be passed into the NewWithClaims function
		UserTokenData: tokenData,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // returns a pointer to a jwt token

	tokenString, err := token.SignedString([]byte(mySecretKey)) // converts the token to a string and returns an error

	if err != nil {
		return ut, err
	}

	ut.UserToken = tokenString
	ut.ExpirationTime = expirationTime
	return ut, nil
}

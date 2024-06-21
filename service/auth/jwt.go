package auth

import (
	"strconv"
	"time"

	"github.com/Ion-Stefan/go-kickstart-backend/config"
	"github.com/golang-jwt/jwt"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	// Define the expiration
	expiration, err := strconv.Atoi(config.Envs.JWTExpiration)
	if err != nil {
		return "", err
	}

	expirationTime := time.Second * time.Duration(expiration)

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    strconv.Itoa(userID),
		"expiredAt": time.Now().Add(expirationTime).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

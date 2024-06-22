package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Ion-Stefan/go-kickstart-backend/config"
	"github.com/Ion-Stefan/go-kickstart-backend/types"
	"github.com/golang-jwt/jwt"
)

type contextKey string

const UserKey contextKey = "userID"

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

func WithJWTAuth(handlerFunc http.HandlerFunc, store types.UserStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the request
		tokenString := getTokenFromRequest(r)

		// validate the JWT
		token, err := validateToken(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}
		if !token.Valid {
			permissionDenied(w)
			return
		}
		// Get the userID from the token
		claims := token.Claims.(jwt.MapClaims)
		userID, err := strconv.Atoi(claims["userID"].(string))
		if err != nil {
			permissionDenied(w)
			return
		}
		u, err := store.GetUserByID(userID)
		if err != nil {
			permissionDenied(w)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserKey, u.ID)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}

func getTokenFromRequest(r *http.Request) string {
	// Get the token from the Authorization header
	token := r.Header.Get("Authorization")
	if token == "" {
		return ""
	}

	return token
}

func validateToken(tokenString string) (*jwt.Token, error) {
	// Parse the token
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Get the secret from the config
		return []byte(config.Envs.JWTSecret), nil
	})
}

func permissionDenied(w http.ResponseWriter) {
	http.Error(w, "Permission denied", http.StatusForbidden)
}

func GetUserIDFromContext(ctx context.Context) int {
	userID := ctx.Value(UserKey)
	if userID == nil {
		return -1
	}

	return userID.(int)
}

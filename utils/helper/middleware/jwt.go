package middleware

import (
	"ecomplaint/model/web"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func ExtractJWTTokenID(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["id"].(string)
	if !ok {
		return "", fmt.Errorf("user id not found in claims")
	}

	return userID, nil
}

func ExtractTokenFromAuthorizationHeader(r *http.Request) (string, error) {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return "", fmt.Errorf("authorization header not found")
	}

	// Mengecek apakah skema adalah "Bearer"
	parts := strings.Split(authorizationHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}

func GenerateToken(userLoginResponse *web.UserLoginResponse, id string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = userLoginResponse.Name
	claims["email"] = userLoginResponse.Email
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

func GenerateTokenUserID(id string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

func GenerateTokenAdminID(id string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_ADMIN")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

func GenerateAdminToken(adminLoginResponse *web.AdminLoginResponse, id string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = adminLoginResponse.Name
	claims["email"] = adminLoginResponse.Email
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_ADMIN")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

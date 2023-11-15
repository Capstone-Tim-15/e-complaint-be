package middleware

import (
	"ecomplaint/model/web"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

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

func GenerateAdminToken(adminLoginResponse *web.AdminLoginResponse, id string) (string, error) {
	expireTime := time.Now().Add(time.Hour * 1).Unix()
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = adminLoginResponse.Name
	claims["email"] = adminLoginResponse.Email
	claims["exp"] = expireTime

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return validToken, nil
}

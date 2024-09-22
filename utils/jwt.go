package utils

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("SECRET_KEY")

func GenerateJWTToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyJwtToken(token string) (float64, string, error) {
	//toekn Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1bWl0MTIzIiwiZXhwIjoxNzI2MzA2MzY3LCJ1c2VySWQiOjR9.siJmcJuzJMro3YZfXhwhT7YBQ0V_8mL9PXhX98OMVRg

	//removing the Bearer from token string
	if len(token) > 6 && token[:7] == "Bearer " {
		token = token[7:]
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signin method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return -1, "", errors.New("failed to parse token")
	}

	tokenIsValid := parsedToken.Valid
	fmt.Println(tokenIsValid)
	if !tokenIsValid {
		return -1, "", errors.New("invalid Token")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return -1, "", errors.New("Invalid token claims")
	}

	user_id := claims["userId"].(float64)
	email := claims["email"].(string)

	return user_id, email, nil
}

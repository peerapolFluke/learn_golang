package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"seamoor/config"
	"time"
)

type Cliams struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func getJwtSecret() string {
	secret := config.Env.JwtSecret
	if secret == "" {
		return "aSecret"
	}
	return secret
}

func JwtGenerate(userID int, email string) (string, error) {
	cliams := &Cliams{
		ID:    userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)

	return jwtToken.SignedString([]byte(getJwtSecret()))
}

func JwtVerify(token string) (*Cliams, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(getJwtSecret()), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Cliams{}, keyFunc)
	if err != nil {
		return nil, err
	}
	payload, ok := jwtToken.Claims.(*Cliams)
	if !ok {
		return nil, fmt.Errorf("Invalid")
	}

	return payload, nil
}

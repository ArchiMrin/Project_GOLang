package utils

import (
	"fmt"
	"time"

	"github.com/ArchiMrin/Project_GOLang/eCOMM/constants"
	"github.com/ArchiMrin/Project_GOLang/eCOMM/entities"
	"github.com/golang-jwt/jwt"
)

func ValidateToken(signedToken string) (claims *entities.SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&entities.SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(constants.SecretKey), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*entities.SignedDetails)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return claims, msg
}

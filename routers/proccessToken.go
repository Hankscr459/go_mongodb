package routers

import (
	"errors"
	"strings"

	"twitter/db"
	"twitter/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var Email string

var IDUser string

func ProccessToken(tk string) (*models.Cliam, bool, string, error) {
	secret := []byte("Thereismysecretkey")
	claims := &models.Cliam{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("Format token invalid")
	}

	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err == nil {
		_, create, _ := db.CheckIsExistUser(claims.Email)
		if create == true {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, create, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("Token invalid")
	}
	return claims, false, string(""), err
}

package jwt

import (
	"time"

	"twitter/models"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerJWT(t models.User) (string, error) {
	secret := []byte("Thereismysecretkey")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"number":    t.Number,
		"name":      t.Name,
		"birth":     t.Birth,
		"biography": t.Biography,
		"location":  t.Location,
		"siteweb":   t.WebSite,
		"_id":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(secret)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}

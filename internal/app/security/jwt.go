package security

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(username, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		subClaim:   username,
		emailClaim: email,
		iatClaim:   time.Now().Unix(),
	})
	return token.SignedString("secret") // TODO signing key from environment variable
}

package security

import (
	"context"
	"github.com/dgrijalva/jwt-go"
)

const (
	userKey    = "user"
	subClaim   = "sub"
	emailClaim = "email"
	iatClaim   = "iat"
)

func Username(ctx context.Context) (string, error) {
	token, ok := ctx.Value(userKey).(*jwt.Token)
	if !ok {
		return "", notAuthenticatedError
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", notAuthenticatedError
	}
	username, ok := claims[subClaim].(string)
	if !ok {
		return "", notAuthenticatedError
	}
	return username, nil
}

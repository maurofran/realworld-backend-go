package security

import "errors"

// Not authenticated error.

var notAuthenticatedError = errors.New("no authentication found in context")

func IsNotAuthenticated(err error) bool {
	return err == notAuthenticatedError
}
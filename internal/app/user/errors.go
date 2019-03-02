package user

import "fmt"

// Not found error.

func notFound(id interface{}) *notFoundError {
	return &notFoundError{id: id}
}

type notFoundError struct {
	id interface{}
}

func (e *notFoundError) Error() string {
	return fmt.Sprintf("no user found for %s", e.id)
}

func (e *notFoundError) UserNotFound() (bool, interface{}) {
	return true, e.id
}

type userNotFounder interface {
	UserNotFound() (bool, interface{})
}

func IsNotFound(err error) (bool, interface{}) {
	if e, ok := err.(userNotFounder); ok {
		return e.UserNotFound()
	}
	return false, nil
}

// Wrong password error

func wrongPassword(id interface{}) *wrongPasswordError {
	return &wrongPasswordError{id: id}
}

type wrongPasswordError struct {
	id interface{}
}

func (e *wrongPasswordError) Error() string {
	return fmt.Sprintf("wrong password provided for user %s", e.id)
}

func (e *wrongPasswordError) WrongPassword() (bool, interface{}) {
	return true, e.id
}

type wrongPassworder interface {
	WrongPassword() (bool, interface{})
}

func IsWrongPassword(err error) (bool, interface{}) {
	if e, ok := err.(wrongPassworder); ok {
		return e.WrongPassword()
	}
	return false, nil
}

// Username exists error

func alreadyExists(username string) *alreadyExistsError {
	return &alreadyExistsError{username: username}
}

type alreadyExistsError struct {
	username string
}

func (e *alreadyExistsError) Error() string {
	return fmt.Sprintf("user with username '%s' already exists", e.username)
}

func (e *alreadyExistsError) AlreadyExists() (bool, string) {
	return true, e.username
}

type alreadyExister interface {
	AlreadyExists() (bool, string)
}

func IsAlreadyExists(err error) (bool, string) {
	if e, ok := err.(alreadyExister); ok {
		return e.AlreadyExists()
	}
	return false, ""
}

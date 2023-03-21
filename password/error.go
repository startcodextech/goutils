package password

import "errors"

var (
	// ErrPasswordNotMatch is returned when the hashed password does not match the plain-text password.
	ErrPasswordNotMatch = errors.New("hashed password does not match plain-text password")
)

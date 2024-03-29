package password

import goerror "github.com/startcodextech/goerrors"

const (
	// ErrPasswordNotMatch is returned when the hashed password does not match the plain-text password.
	ErrPasswordNotMatch = goerror.Error("hashed password does not match plain-text password")
)

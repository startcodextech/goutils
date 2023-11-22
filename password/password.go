package password

import (
	"golang.org/x/crypto/bcrypt"
)

type (
	// Password is a struct that holds the hashed password.
	Password struct {
		hashedPassword string
	}
)

// NewPassword returns a noop password.
func NoopPassword() Password {
	return Password{
		hashedPassword: "",
	}
}

// HashPassword returns a new password from plain password text.
func HashPassword(password string, cost int) (Password, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return NoopPassword(), err
	}
	return Password{
		hashedPassword: string(hashed),
	}, nil
}

func HashPasswordString(password string, cost int) string {
	pass, _ := HashPassword(password, cost)

	return pass.String()
}

// FromHash creates a Password from a hashed password and its corresponding plain-text password.
// It returns an error if the hashed password does not match the plain-text password.
func FromHash(hashedPassword string, password string) (Password, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return NoopPassword(), ErrPasswordNotMatch
	}
	return Password{
		hashedPassword: hashedPassword,
	}, nil
}

// IsNil returns true if the password is nil or empty.
func (p *Password) IsNil() bool {
	return p.hashedPassword == ""
}

// String returns the hashed password.
func (p *Password) String() string {
	return p.hashedPassword
}

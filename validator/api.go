package validator

import (
	"regexp"
	"unicode"
)

const (
	NameOrLastnameRegexp = `^[a-zA-ZáéíóúÁÉÍÓÚñÑüÜ' -]{2,50}$`
	PasswordRegexp       = `[@[\]^_!"#$%&'()*+,-./:;{}<>|=~?]`
)

func Regexp(value, regex string) bool {
	regx := regexp.MustCompile(regex)

	return regx.MatchString(value)
}

func IsNameOrLastname(value string) bool {
	return Regexp(value, NameOrLastnameRegexp)
}

func IsEmail(value string) bool {
	return Regexp(value, EmailRegexp)
}

// IsValidPassword validate password with the following criteria
//
// * They must contain both uppercase and lowercase letters.
//
// * They must contain numbers and special signs.
//
// * They must be a minimum of 8 characters.
//
// * Specific list of allowed special characters: @ [ ] ^ _ ! " # $ % & ' ( ) * + , - . / : ; { } < > = | ~ ?
func IsValidPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	var hasUpper, hasLower, hasNumber, hasSpecial bool

	var specialCharRegex = regexp.MustCompile(PasswordRegexp)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case specialCharRegex.MatchString(string(char)):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial
}

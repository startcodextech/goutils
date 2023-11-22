package validator

import "regexp"

const (
	NameOrLastnameRegexp = `^[a-zA-ZáéíóúÁÉÍÓÚñÑüÜ' -]{2,50}$`
	PasswordRegexp       = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@[\]^_!"#$%&'()*+,-./:;{}<>|=~?])[A-Za-z\d@[\]^_!"#$%&'()*+,-./:;{}<>|=~?]{8,}$`
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
func IsValidPassword(value string) bool {
	return Regexp(value, PasswordRegexp)
}

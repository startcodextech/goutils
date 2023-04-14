package validator

import "regexp"

type Email string

func (e Email) IsValid() bool {
	expresionRegular := `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	valido, err := regexp.MatchString(expresionRegular, string(e))
	return valido && err == nil
}

func (e Email) String() string {
	return string(e)
}

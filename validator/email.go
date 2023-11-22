package validator

const (
	EmailRegexp = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
)

type Email string

func (e Email) IsValid() bool {
	return Regexp(string(e), EmailRegexp)
}

func (e Email) String() string {
	return string(e)
}

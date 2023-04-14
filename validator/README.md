# Validator

It is a package that provides functionality to perform simple validation

## import

```go
import (
    "github.com/start-codex/utils/validator"
)
```

## Validate Email

```go
email := validator.Email("test@user.com")

if email.IsValid() {
    fmt.Println("Email is valid")
} else {
    fmt.Println("Email is invalid")
}
```
# Package Password

This package has as objective functions for the validation and generation of passwords.

## Import

```go
import "github.com/developersgotech/go-util/password"
```

## Usage

### Generate a new password hash

```go
hashed, err := password.HashPassword("my-password", 14)
```

### Create password from hash existing
    
```go
hashed, err := password.FromHash("$2a$14$2b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0u1v2w3x4y5z6", "my-password")
```

### Get value of password hash string

```go
hash := hashed.String()
```
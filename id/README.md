# Package Id

This package has as objective functions for the validation and generation of uuid v4.

## Import

```go
import "github.com/start-codex/goutils/id"
```

## Usage

### Generate a new uuid v4
```go
uuid, err := id.New()
```

### Parse a uuid v4 from string
```go
valid, err := id.Parse("a1b2c3d4-e5f6-g7h8-i9j0-k1l2m3n4o5p6")
```

### Validate is zero value `00000000-0000-0000-0000-000000000000`
```go
isZero := id.IsZero()
```

### Get value to string
```go
value := id.ID()
```

### Generate value to string
```go
value := id.NewString()
```

## Deprecation

### String()
The function `String()` is deprecated, use `ID()` instead.
```go
uuid, err := id.New()

if err != nil {
    panic(err)
}
	
// uuid.String() replace with
value := uuid.ID()
```
package id

import (
	"github.com/google/uuid"
)

type (
	// ID is a unique identifier.
	ID struct {
		value uuid.UUID
	}
)

// NoopId returns a noop ID.
func NoopId() ID {
	return ID{value: uuid.UUID{}}
}

// NewID returns a new ID.
func NewID() (ID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return NoopId(), err
	}
	return ID{value: id}, nil
}

// ParseID returns an ID from a string.
func ParseID(value string) (ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return NoopId(), err
	}
	return ID{value: id}, nil
}

// IsZero returns true if the ID is nil.
func (id ID) IsZero() bool {
	return id.value.String() == ID_NIL_VALUE
}

// String returns the string representation of the ID.
func (id ID) String() string {
	return id.value.String()
}

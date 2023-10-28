package id

import (
	"github.com/google/uuid"
)

type (
	// IDer is an interface that returns an ID.
	IDer interface {
		ID() string
	}

	// ID is a unique identifier.
	ID struct {
		value uuid.UUID
	}
)

// NoopId returns a noop ID.
func NoopId() ID {
	return ID{value: uuid.UUID{}}
}

// New returns a new ID.
func New() (ID, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return NoopId(), err
	}
	return ID{value: id}, nil
}

// NewString returns a new ID as a string.
func NewString() string {
	id, _ := New()
	return id.String()
}

// ParseID returns an ID from a string.
func Parse(value string) (ID, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return NoopId(), err
	}
	return ID{value: id}, nil
}

// ID returns the string representation of the ID.
func (id ID) ID() string {
	return id.value.String()
}

// IsZero returns true if the ID is nil.
func (id ID) IsZero() bool {
	return id.value.String() == ID_NIL_VALUE
}

// Deprecated: use IDer.ID() instead.
func (id ID) String() string {
	return id.value.String()
}

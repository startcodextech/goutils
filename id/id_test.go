package id

import "testing"

func TestNewID(t *testing.T) {
	id, err := NewID()
	if err != nil {
		t.Errorf("error creating new ID: %v", err)
	}
	if id.IsZero() {
		t.Errorf("new ID should not be zero")
	}
	_, err = ParseID(NoopId().String())
	if err != nil {
		t.Errorf("error parsing noop ID: %v", err)
	}
	if id.String() == ID_NIL_VALUE {
		t.Errorf("new ID should not be nil")
	}
}

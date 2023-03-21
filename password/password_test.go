package password

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	pass, err := HashPassword("password", 14)
	if err != nil {
		t.Errorf("error hashing password: %v", err)
	}
	if pass.IsNil() {
		t.Errorf("password should not be nil")
	}
	t.Log("password hashed", pass.String())
}

func TestFromHash(t *testing.T) {
	_, err := FromHash("$2a$14$HIGCGqQxfpaiZiTGfeRcmOKWLmBIvj1kG6.HZPUzvMwYmZDNFjGsG", "password")
	if err != nil {
		t.Errorf("error creating password from hash: %v", err)
	}
}

func TestNoopPassword(t *testing.T) {
	pass := NoopPassword()
	if !pass.IsNil() {
		t.Errorf("password should be nil")
	}
}

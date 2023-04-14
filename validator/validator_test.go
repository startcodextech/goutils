package validator

import "testing"

func TestEmail_IsValid(t *testing.T) {
	email := Email("test@user.com")

	if !email.IsValid() {
		t.Error("Email is not valid")
	}

	t.Log("Email_IsValid: OK")
}

package utils

import "testing"

func TestPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Error("Error hashing password")
	}
	if err = CheckPasswordHash(password, hash); err != nil {
		t.Error("Error checking password hash")
	}

	if err = CheckPasswordHash("wrongpassword", hash); err == nil {
		t.Error("Error checking password hash")
	}
}

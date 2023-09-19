package utils

import "testing"

func TestPassword(t *testing.T) {
	password := "password"
	hash, err := HashPassword(password)
	if err != nil {
		t.Error("Error hashing password")
	}
	if !CheckPasswordHash(password, hash) {
		t.Error("Error checking password hash")
	}

	if CheckPasswordHash("wrongpassword", hash) {
		t.Error("Error checking password hash")
	}
}

package auth

import "testing"

func TestHashAndCheckPassword(t *testing.T) {
	hash, err := HashPassword("secret123")
	if err != nil {
		t.Fatalf("hashing failed: %v", err)
	}

	if !CheckPasswordHash("secret123", hash) {
		t.Errorf("expected correct password to match")
	}
}
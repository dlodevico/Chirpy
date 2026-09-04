package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashAndCheckPassword(t *testing.T) {
	password := "mysecretpassword"
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error hashing password, got %v", err)
	}

	match, err := CheckPasswordHash(password, hash)
	if err != nil {
		t.Fatalf("expected no error checking password hash, got %v", err)
	}
	if !match {
		t.Fatalf("expected password to match hash")
	}

	match, err = CheckPasswordHash("wrongpassword", hash)
	if err != nil {
		t.Fatalf("expected no error checking wrong password hash, got %v", err)
	}
	if match {
		t.Fatalf("expected wrong password to not match hash")
	}
}

func TestMakeAndValidateJWT(t *testing.T) {
	userID := uuid.New()
	secret := "supersecretkey"

	tokenString, err := MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("expected no error making JWT, got %v", err)
	}

	extractedID, err := ValidateJWT(tokenString, secret)
	if err != nil {
		t.Fatalf("expected no error validating JWT, got %v", err)
	}
	if extractedID != userID {
		t.Fatalf("expected user ID %v, got %v", userID, extractedID)
	}

	expiredTokenString, err := MakeJWT(userID, secret, -time.Hour)
	if err != nil {
		t.Fatalf("expected no error making expired JWT, got %v", err)
	}

	_, err = ValidateJWT(expiredTokenString, secret)
	if err == nil {
		t.Fatalf("expected error validating expired JWT, got nil")
	}

	wrongSecret := "differentsecretkey"
	_, err = ValidateJWT(tokenString, wrongSecret)
	if err == nil {
		t.Fatalf("expected error validating JWT with wrong secret, got nil")
	}
}

func TestGetBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer my-secret-token")

	token, err := GetBearerToken(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "my-secret-token" {
		t.Fatalf("expected 'my-secret-token', got %s", token)
	}

	emptyHeaders := http.Header{}
	_, err = GetBearerToken(emptyHeaders)
	if err == nil {
		t.Fatalf("expected error for missing authorization header, got nil")
	}

	badHeaders := http.Header{}
	badHeaders.Set("Authorization", "Basic my-secret-token")
	_, err = GetBearerToken(badHeaders)
	if err == nil {
		t.Fatalf("expected error for malformed authorization header, got nil")
	}
}

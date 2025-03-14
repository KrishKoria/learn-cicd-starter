package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey api-key-12345")

	apiKey, err := GetAPIKey(headers)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	expectedKey := "api-key-12345"
	if apiKey != expectedKey {
		t.Errorf("Expected API key to be %s, got %s", expectedKey, apiKey)
	}
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
	headers := http.Header{}

	_, err := GetAPIKey(headers)

	if err != ErrNoAuthHeaderIncluded {
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKey_InvalidFormat(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "InvalidFormat")

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Error("Expected an error for invalid Authorization format, got nil")
	}
}

func TestGetAPIKey_WrongPrefix(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer api-key-12345")

	_, err := GetAPIKey(headers)

	if err == nil {
		t.Error("Expected an error for wrong prefix, got nil")
	}
}

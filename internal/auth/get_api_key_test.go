package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name        string
		headers     http.Header
		expected    string
		expectError bool
	}{
		{
			name: "Valid API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey my-secret-api-key"},
			},
			expected:    "my-secret-api-key",
			expectError: false,
		},
		{
			name:        "Missing Authorization Header",
			headers:     http.Header{},
			expected:    "",
			expectError: true,
		},
		{
			name: "Malformed Header - Wrong Prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer my-secret-api-key"},
			},
			expected:    "",
			expectError: true,
		},
		{
			name: "Malformed Header - Missing Key Value",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual, err := GetAPIKey(tt.headers)

			if (err != nil) != tt.expectError {
				t.Errorf("Test %q: expected error = %v, got error = %v", tt.name, tt.expectError, err)
			}

			if actual != tt.expected {
				t.Errorf("Test %q: expected key = %q, got key = %q", tt.name, tt.expected, actual)
			}
		})
	}
}

package auth

import (
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		header        http.Header
		expectedKey   string
		expectedError bool
	}{
		"Valid Header": {
			header:        http.Header{"Authorization": []string{"ApiKey secret123"}},
			expectedKey:   "secret123",
			expectedError: false,
		},
		"Malformed Header": {
			header:        http.Header{"Authorization": []string{"secret123"}},
			expectedKey:   "",
			expectedError: true,
		},
		"Missing Authorization Key": {
			header:        http.Header{"": []string{"ApiKey secret123"}},
			expectedKey:   "",
			expectedError: true,
		},
		"Empty Header": {
			header:        http.Header{},
			expectedKey:   "",
			expectedError: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			key, err := GetAPIKey(tc.header)
			if (err != nil) != tc.expectedError {
				t.Fatalf("expected error: %v, got: %v", tc.expectedError, err)
			}

			diff := cmp.Diff(tc.expectedKey, key)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

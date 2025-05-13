package auth

import (
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	testcases := []struct {
		name    string
		headers http.Header
		wantErr bool
	}{
		{
			name: "valid headers",
			headers: http.Header{
				"Authorization": []string{"ApiKey validApiKey"},
			},
			wantErr: false,
		},
		{
			name: "Only prefix",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			wantErr: true,
		},
		{
			name: "Wrong prefix",
			headers: http.Header{
				"Authorization": []string{"Bearer"},
			},
			wantErr: true,
		},
		{
			name:    "No Auth Header",
			headers: http.Header{},
			wantErr: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := GetAPIKey(tc.headers)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetAPIKey error = %q, wantErr = %v", err, tc.wantErr)
			}
		})
	}
}

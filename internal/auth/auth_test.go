package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		input   string
		want    string
		wantErr bool
	}{
		"valid key": {
			input:   "ApiKey mykey",
			want:    "mykey",
			wantErr: false,
		},
		"missing key": {
			input:   "gibberish",
			want:    "",
			wantErr: true,
		},
		"extra whitespace": {
			input:   " ApiKey  mykey",
			want:    "",
			wantErr: true,
		},
		"empty header": {
			wantErr: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			headers := http.Header{}
			headers.Set("Authorization", tt.input)

			got, err := GetAPIKey(headers)
			if err != nil && !tt.wantErr {
				t.Errorf("Expected error for %#v but got %#v", tt.input, err)
			}
			if got != tt.want {
				t.Errorf("expected: %#v, but got %#v", tt.want, got)
			}
		})
	}
}

//ff:func feature=validate type=test control=iteration dimension=1 topic=response-body-required
//ff:what isErrorStatus — 4xx/5xx 판별 + 204/304 예외 + 비숫자 검증

package openapi

import "testing"

func TestIsErrorStatus(t *testing.T) {
	tests := []struct {
		status string
		want   bool
	}{
		{"200", false},
		{"201", false},
		{"204", false},  // explicitly bodyless
		{"301", false},
		{"304", false},  // explicitly bodyless
		{"400", true},
		{"401", true},
		{"404", true},
		{"422", true},
		{"499", true},
		{"500", true},
		{"503", true},
		{"599", true},
		{"600", false},
		{"default", false},
		{"2XX", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.status, func(t *testing.T) {
			got := isErrorStatus(tt.status)
			if got != tt.want {
				t.Errorf("isErrorStatus(%q) = %v, want %v", tt.status, got, tt.want)
			}
		})
	}
}

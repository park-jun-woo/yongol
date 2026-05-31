//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what TestHurlOpenAPIHelpers — unit tests for the pure hurl_openapi helper functions
package hurl_openapi

import (
	"testing"
)

func TestXoh13GuardDefaultStatus(t *testing.T) {
	tests := []struct {
		ty   string
		want int
	}{
		{"empty", 404},
		{"exists", 409},
		{"auth", 403},
		{"state", 409},
		{"eval", 0},
		{"get", 0},
	}
	for _, tt := range tests {
		if got := xoh13GuardDefaultStatus(tt.ty); got != tt.want {
			t.Errorf("xoh13GuardDefaultStatus(%q) = %d, want %d", tt.ty, got, tt.want)
		}
	}
}

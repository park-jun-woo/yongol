//ff:func feature=external type=test control=iteration dimension=1
//ff:what TestCaseHelpers — lcFirst/toCamelCase/toPascalCase 변환 검증
package external

import (
	"testing"
)

func TestToCamelCase(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"UserName", "userName"},
		{"user_name", "userName"},
		{"ID", "id"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := toCamelCase(tt.in); got != tt.want {
			t.Errorf("toCamelCase(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

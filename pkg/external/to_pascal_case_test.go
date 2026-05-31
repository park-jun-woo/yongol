//ff:func feature=external type=test control=iteration dimension=1
//ff:what TestCaseHelpers — lcFirst/toCamelCase/toPascalCase 변환 검증
package external

import (
	"testing"
)

func TestToPascalCase(t *testing.T) {
	tests := []struct {
		in, want string
	}{
		{"user_name", "UserName"},
		{"userName", "UserName"},
		{"foo-bar", "FooBar"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := toPascalCase(tt.in); got != tt.want {
			t.Errorf("toPascalCase(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

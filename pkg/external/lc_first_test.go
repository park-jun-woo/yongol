//ff:func feature=external type=test control=iteration dimension=1
//ff:what TestCaseHelpers — lcFirst/toCamelCase/toPascalCase 변환 검증
package external

import (
	"testing"
)

func TestLcFirst(t *testing.T) {
	// lcFirst delegates to the same Go camelCase conversion as toCamelCase.
	tests := []struct {
		in, want string
	}{
		{"UserName", "userName"},
		{"FooBar", "fooBar"},
		{"", ""},
	}
	for _, tt := range tests {
		if got := lcFirst(tt.in); got != tt.want {
			t.Errorf("lcFirst(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

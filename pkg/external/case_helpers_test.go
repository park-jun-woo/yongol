//ff:func feature=external type=test control=sequence
//ff:what TestCaseHelpers — lcFirst/toCamelCase/toPascalCase 변환 검증

package external

import "testing"

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

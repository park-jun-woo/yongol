//ff:func feature=validate type=test control=iteration dimension=1 topic=ddl-structural
//ff:what TestDDLHelpers — unit tests for the pure DDL validate helper functions
package ddl

import (
	"testing"
)

func TestHasInlineSensitiveAnnotation(t *testing.T) {
	tests := []struct {
		line string
		want bool
	}{
		{"password TEXT -- @sensitive", true},
		{"password TEXT --@sensitive", true},
		{"col TEXT -- @SENSITIVE", true},
		{"col TEXT -- @nosensitive", true},
		{"col TEXT --@nosensitive", true},
		{"password TEXT", false},
		{"col TEXT -- some comment", false},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got := hasInlineSensitiveAnnotation(tt.line); got != tt.want {
				t.Errorf("hasInlineSensitiveAnnotation(%q) = %v, want %v", tt.line, got, tt.want)
			}
		})
	}
}

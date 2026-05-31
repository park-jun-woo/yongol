//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestParseInputValue -- parseInputValue 숫자/따옴표/변수/dotted 리터럴 분류 검증
package ir

import (
	"testing"
)

func TestIsNumericLiteral(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"1", true},
		{"42", true},
		{"-3", true},
		{"1.5", true},
		{"-0.25", true},
		{"0", true},
		{"", false},
		{"-", false},
		{"abc", false},
		{"1.2.3", false},
		{"wf.ID", false},
		{"request.email", false},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := isNumericLiteral(tt.input)
			if got != tt.want {
				t.Errorf("isNumericLiteral(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

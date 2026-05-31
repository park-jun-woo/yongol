//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-statemachine
//ff:what TestInferStateLiteralType — inferStateLiteralType 리터럴 타입 추론 분기 검증
package ssac_statemachine

import (
	"testing"
)

func TestInferStateLiteralType(t *testing.T) {
	cases := []struct {
		value string
		want  string
	}{
		{`"hello"`, "string"},
		{"true", "bool"},
		{"false", "bool"},
		{"nil", "nil"},
		{"42", "int"},
		{"-7", "int"},
		{"3.14", "float64"},
		{"-2.5", "float64"},
		{"abc", ""},
		{"", ""},
		{`"`, ""},
	}
	for _, tc := range cases {
		t.Run(tc.value, func(t *testing.T) {
			if got := inferStateLiteralType(tc.value); got != tc.want {
				t.Errorf("inferStateLiteralType(%q) = %q, want %q", tc.value, got, tc.want)
			}
		})
	}
}

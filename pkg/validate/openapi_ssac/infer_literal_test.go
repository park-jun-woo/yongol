//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ssac
//ff:what inferLiteral — string/bool/nil/int64/float64/unknown 타입 추론 검증

package openapi_ssac

import "testing"

func TestInferLiteral(t *testing.T) {
	tests := []struct {
		value string
		want  string
	}{
		{`"hello"`, "string"},
		{`""`, "string"},
		{"true", "bool"},
		{"false", "bool"},
		{"nil", "nil"},
		{"42", "int64"},
		{"-1", "int64"},
		{"3.14", "float64"},
		{"someVar", ""},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			got := inferLiteral(tt.value)
			if got != tt.want {
				t.Errorf("inferLiteral(%q) = %q, want %q", tt.value, got, tt.want)
			}
		})
	}
}

//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ssac
//ff:what normalizeTypeName — slice/pointer/wrapper/pkg prefix 제거 검증

package openapi_ssac

import "testing"

func TestNormalizeTypeName(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"User", "User"},
		{"*User", "User"},
		{"[]User", "User"},
		{"[]billing.CheckCreditsResponse", "CheckCreditsResponse"},
		{"Page[Workflow]", "Workflow"},
		{"Cursor[User]", "User"},
		{"*pkg.Type", "Type"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := normalizeTypeName(tt.input)
			if got != tt.want {
				t.Errorf("normalizeTypeName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

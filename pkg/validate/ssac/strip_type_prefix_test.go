//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what stripTypePrefix 단위 테스트

package ssac

import "testing"

func TestStripTypePrefix(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"Course", "Course"},
		{"[]Course", "Course"},
		{"billing.CheckResp", "CheckResp"},
		{"[]billing.CheckResp", "CheckResp"},
		{"", ""},
	}
	for _, tt := range tests {
		got := stripTypePrefix(tt.input)
		if got != tt.want {
			t.Errorf("stripTypePrefix(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

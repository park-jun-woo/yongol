//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what splitEvalModel 단위 테스트

package ssac

import "testing"

func TestSplitEvalModel(t *testing.T) {
	tests := []struct {
		input    string
		wantPkg  string
		wantFunc string
	}{
		{"billing.IsZero", "billing", "IsZero"},
		{"pkg.Func", "pkg", "Func"},
		{"nodot", "", ""},
		{"", "", ""},
		{".Leading", "", ""},
	}
	for _, tt := range tests {
		pkg, fn := splitEvalModel(tt.input)
		if pkg != tt.wantPkg || fn != tt.wantFunc {
			t.Errorf("splitEvalModel(%q) = (%q, %q), want (%q, %q)", tt.input, pkg, fn, tt.wantPkg, tt.wantFunc)
		}
	}
}

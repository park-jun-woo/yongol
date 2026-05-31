//ff:func feature=ssac-parse type=test control=iteration dimension=1
//ff:what TestSSaCParseHelpers — unit tests for the pure ssac parser helper functions
package ssac

import (
	"testing"
)

func TestSplitPackagePrefix(t *testing.T) {
	tests := []struct {
		in       string
		wantPkg  string
		wantRest string
	}{
		{"session.Session.Get", "session", "Session.Get"},
		{"Course.FindByID", "", "Course.FindByID"},   // 2-part, no pkg
		{"plain", "", "plain"},                       // no dot
		{"Pkg.Model.Method", "", "Pkg.Model.Method"}, // uppercase first → no pkg
	}
	for _, tt := range tests {
		pkg, rest := splitPackagePrefix(tt.in)
		if pkg != tt.wantPkg || rest != tt.wantRest {
			t.Errorf("splitPackagePrefix(%q) = (%q,%q), want (%q,%q)", tt.in, pkg, rest, tt.wantPkg, tt.wantRest)
		}
	}
}

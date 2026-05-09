//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-73 — SSaC import는 full Go import path 필수 (bare name 거부)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s73ImportMustBeFullPath validates S-73: every SSaC import path must be a
// full Go import path containing at least one '/'. Bare names like "dashboard"
// are rejected — the user must write the complete path (e.g.
// "github.com/user/project/internal/dashboard").
func s73ImportMustBeFullPath(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, imp := range fn.Imports {
			if strings.Contains(imp, "/") {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    fn.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-73] SSaC import %q is a bare name — use full Go import path", imp),
				Advice:  fmt.Sprintf("Replace with the full path, e.g. \"github.com/.../internal/%s\".", imp),
			})
		}
	}
	return diags
}

//ff:func feature=validate type=rule control=sequence topic=hurl-structural
//ff:what H-2 — WARNING when tests/ directory exists but contains no .hurl files

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// h02EmptyTestsDir validates H-2: when tests/ is declared (SSOTDeclared) but
// contains no .hurl scenario files, emit a WARNING. Distinguishes "user opted
// out of scenarios" (SSOTAbsent) from "user declared tests dir but left it
// empty" which is almost always WIP or mistake.
func h02EmptyTestsDir(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.PresenceOf(yongol.KindScenario) != yongol.SSOTDeclared {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "tests/",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[H-2] tests/ directory exists but contains no .hurl scenario files",
		Advice:  "Add scenario or invariant Hurl files, or remove the tests/ directory",
	}}
}

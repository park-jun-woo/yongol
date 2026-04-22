//ff:func feature=validate type=rule control=sequence topic=hurl-structural
//ff:what H-2 — tests/ 디렉토리 있는데 .hurl 파일 0 개면 WARNING

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
		Advice:  "scenario/invariant hurl 파일을 추가하거나 tests/ 디렉토리를 제거하세요",
	}}
}

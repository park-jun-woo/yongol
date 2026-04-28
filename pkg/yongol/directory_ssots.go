//ff:func feature=orchestrator type=loader control=sequence
//ff:what directorySSOTs — specs abs 루트 아래 디렉토리 기반 SSOT 후보 목록 반환
package yongol

import (
	"path/filepath"
)

// directorySSOTs returns the directory-backed SSOT candidates (everything
// except single-file SSOTs like manifest.yaml and api/openapi.yaml) keyed to
// the given absolute specs root.
func directorySSOTs(abs string) []dirSSOT {
	return []dirSSOT{
		{KindDDL, filepath.Join(abs, "db"), []string{"*.sql"}},
		{KindSSaC, filepath.Join(abs, "service"), []string{"*.ssac", "*/*.ssac"}},
		{KindStates, filepath.Join(abs, "states"), []string{"*.md"}},
		{KindPolicy, filepath.Join(abs, "policy"), []string{"*.rego"}},
		{KindScenario, filepath.Join(abs, "tests"), []string{"smoke.hurl", "scenario-*.hurl", "invariant-*.hurl"}},
		{KindFunc, filepath.Join(abs, "func"), []string{"*/*.go"}},
		{KindTSX, filepath.Join(abs, "frontend"), []string{"*.tsx", "*/*.tsx", "*/*/*.tsx", "*/*/*/*.tsx"}},
	}
}

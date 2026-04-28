//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what DetectSSOTs — tests/smoke.hurl 단독·혼합·부재 분기 표기반 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

// TestDirectorySSOTs_Smoke verifies that the KindScenario glob recognises a
// bare `smoke.hurl` file (PhaseV001 made it user-authored). Three rows cover
// the matrix: smoke-only, smoke+scenario+invariant mixed, and an empty
// `tests/` directory (declared-but-unpopulated).
func TestDirectorySSOTs_Smoke(t *testing.T) {
	cases := []struct {
		name         string
		files        []string
		wantPresence SSOTPresence
	}{
		{"smoke-only", []string{"smoke.hurl"}, SSOTPopulated},
		{"mixed", []string{"smoke.hurl", "scenario-foo.hurl", "invariant-bar.hurl"}, SSOTPopulated},
		{"none", nil, SSOTDeclared},
	}

	for _, tc := range cases {
		tmp := newTmpSpecsDir(t)
		dir := filepath.Join(tmp, "tests")
		if err := os.MkdirAll(dir, 0o755); err != nil {
			t.Fatalf("%s: mkdir %s: %v", tc.name, dir, err)
		}
		for _, name := range tc.files {
			writeFile(t, filepath.Join(dir, name), "GET {{host}}/ping\nHTTP 200\n")
		}
		detected, err := DetectSSOTs(tmp)
		if err != nil {
			t.Fatalf("%s: DetectSSOTs: %v", tc.name, err)
		}
		d, ok := hasKind(detected, KindScenario)
		if !ok {
			t.Fatalf("%s: KindScenario not detected; detected=%+v", tc.name, detected)
		}
		if d.Presence != tc.wantPresence {
			t.Fatalf("%s: presence = %v; want %v", tc.name, d.Presence, tc.wantPresence)
		}
	}
}

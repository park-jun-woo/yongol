//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what TestDetectDirSSOT/directorySSOTs — glob 매칭 presence 결정 및 후보 목록 검증
package yongol

import (
	"path/filepath"
	"testing"
)

func TestDirectorySSOTs(t *testing.T) {
	abs := "/specs/root"
	got := directorySSOTs(abs)

	want := map[SSOTKind]string{
		KindDDL:      filepath.Join(abs, "db"),
		KindSSaC:     filepath.Join(abs, "service"),
		KindStates:   filepath.Join(abs, "states"),
		KindPolicy:   filepath.Join(abs, "policy"),
		KindScenario: filepath.Join(abs, "tests"),
		KindFunc:     filepath.Join(abs, "func"),
		KindSTML:     filepath.Join(abs, "frontend"),
	}
	if len(got) != len(want) {
		t.Fatalf("expected %d candidates, got %d", len(want), len(got))
	}
	for _, d := range got {
		wantDir, ok := want[d.kind]
		if !ok {
			t.Errorf("unexpected kind %q", d.kind)
			continue
		}
		if d.dir != wantDir {
			t.Errorf("kind %q dir = %q, want %q", d.kind, d.dir, wantDir)
		}
		if len(d.globs) == 0 {
			t.Errorf("kind %q has no globs", d.kind)
		}
	}
}

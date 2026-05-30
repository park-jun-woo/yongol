//ff:func feature=orchestrator type=test control=sequence
//ff:what TestDetectDirSSOT/directorySSOTs — glob 매칭 presence 결정 및 후보 목록 검증

package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectDirSSOTPopulated(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "schema.sql"), []byte("--"), 0o644); err != nil {
		t.Fatal(err)
	}
	d := dirSSOT{kind: KindDDL, dir: dir, globs: []string{"*.sql"}}

	got, err := detectDirSSOT(d)
	if err != nil {
		t.Fatal(err)
	}
	if got.Presence != SSOTPopulated {
		t.Errorf("Presence = %v, want SSOTPopulated", got.Presence)
	}
	if got.Kind != KindDDL || got.Path != dir {
		t.Errorf("got %+v", got)
	}
}

func TestDetectDirSSOTDeclared(t *testing.T) {
	dir := t.TempDir() // exists but no *.sql
	d := dirSSOT{kind: KindDDL, dir: dir, globs: []string{"*.sql"}}

	got, err := detectDirSSOT(d)
	if err != nil {
		t.Fatal(err)
	}
	if got.Presence != SSOTDeclared {
		t.Errorf("Presence = %v, want SSOTDeclared", got.Presence)
	}
	if got.Kind != KindDDL {
		t.Errorf("Kind = %v, want KindDDL", got.Kind)
	}
}

func TestDetectDirSSOTAbsent(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-such-dir")
	d := dirSSOT{kind: KindDDL, dir: missing, globs: []string{"*.sql"}}

	got, err := detectDirSSOT(d)
	if err != nil {
		t.Fatal(err)
	}
	if got.Presence != SSOTAbsent {
		t.Errorf("Presence = %v, want SSOTAbsent", got.Presence)
	}
	// Absent result carries no Kind/Path.
	if got.Kind != "" {
		t.Errorf("expected empty Kind for absent, got %q", got.Kind)
	}
}

func TestDetectDirSSOTGlobError(t *testing.T) {
	// A malformed glob pattern ("[") triggers filepath.ErrBadPattern,
	// exercising the hard-error branch.
	d := dirSSOT{kind: KindDDL, dir: t.TempDir(), globs: []string{"["}}
	_, err := detectDirSSOT(d)
	if err == nil {
		t.Fatal("expected error for malformed glob pattern")
	}
}

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

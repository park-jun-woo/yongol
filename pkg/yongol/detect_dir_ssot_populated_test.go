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

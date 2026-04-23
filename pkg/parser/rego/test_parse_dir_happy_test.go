//ff:func feature=policy type=test control=sequence
//ff:what ParseDir — 정상 rego 파일 1개 → modules 1개

package rego

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseDir_Happy(t *testing.T) {
	dir := t.TempDir()
	content := "package authz\n\ndefault allow := false\n"
	if err := os.WriteFile(filepath.Join(dir, "a.rego"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	mods, diags := ParseDir(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(mods) != 1 {
		t.Fatalf("modules count = %d, want 1", len(mods))
	}
}

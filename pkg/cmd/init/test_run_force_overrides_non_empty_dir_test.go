//ff:func feature=cli-init type=test control=sequence
//ff:what TestRunForceOverridesNonEmptyDir — Run with --force writes even into a non-empty directory

package cliinit

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRunForceOverridesNonEmptyDir(t *testing.T) {
	tmp := t.TempDir()
	if err := os.WriteFile(filepath.Join(tmp, "existing.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:   "Myapp",
		Description: "Test",
		Dir:         tmp,
		Module:      "github.com/test/myapp",
		Force:       true,
	})
	if err != nil {
		t.Fatalf("Run with --force unexpectedly failed: %v", err)
	}
	if _, err := os.Stat(filepath.Join(tmp, "specs", "manifest.yaml")); err != nil {
		t.Errorf("manifest missing after --force run: %v", err)
	}
}

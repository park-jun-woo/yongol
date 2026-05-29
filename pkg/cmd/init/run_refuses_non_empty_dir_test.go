//ff:func feature=cli-init type=test control=sequence
//ff:what TestRunRefusesNonEmptyDir — Run refuses to write into a non-empty directory without --force

package cliinit

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunRefusesNonEmptyDir(t *testing.T) {
	featPath := writeTempFeatures(t)
	tmp := t.TempDir()
	// Pre-populate the target so Run must refuse.
	if err := os.WriteFile(filepath.Join(tmp, "existing.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:    "Myapp",
		FeaturesPath: featPath,
		Description:  "Test",
		Dir:          tmp,
		Module:       "github.com/test/myapp",
	})
	if err == nil {
		t.Fatalf("Run should refuse non-empty dir without --force")
	}
	if !strings.Contains(err.Error(), "not empty") {
		t.Errorf("error message does not mention emptiness: %v", err)
	}
}

//ff:func feature=cli-init type=test control=sequence
//ff:what TestRunDescriptionDefaultsFromProjectID

package cliinit

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunDescriptionDefaultsFromProjectID(t *testing.T) {
	featPath := writeTempFeatures(t)
	tmp := t.TempDir()
	target := filepath.Join(tmp, "myapp")
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID:    "Myapp",
		FeaturesPath: featPath,
		// Description intentionally left empty
		Dir:    target,
		Module: "github.com/test/myapp",
	})
	if err != nil {
		t.Fatalf("Run returned error: %v", err)
	}
	// Should not fail even without description.
	manifest, _ := os.ReadFile(filepath.Join(target, "specs", "manifest.yaml"))
	if !strings.Contains(string(manifest), "Myapp project") {
		t.Errorf("manifest should use default description, got: %s", string(manifest))
	}
}

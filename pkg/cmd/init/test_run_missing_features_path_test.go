//ff:func feature=cli-init type=test control=sequence
//ff:what TestRunMissingFeaturesPath

package cliinit

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunMissingFeaturesPath(t *testing.T) {
	var outBuf, errBuf bytes.Buffer
	err := Run(&outBuf, &errBuf, Options{
		ProjectID: "Myapp",
		// FeaturesPath intentionally empty
		Description: "Test",
		Dir:         t.TempDir(),
		Module:      "github.com/test/myapp",
	})
	if err == nil {
		t.Fatal("Run should fail without features path")
	}
	if !strings.Contains(err.Error(), "features.yaml path is required") {
		t.Errorf("unexpected error: %v", err)
	}
}

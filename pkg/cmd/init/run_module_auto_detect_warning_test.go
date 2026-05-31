//ff:func feature=cli-init type=test control=sequence
//ff:what TestRunExtraBranches — loadFeatures 에러 / Dir 기본값 / Module 자동탐지 warning 분기 검증
package cliinit

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun_ModuleAutoDetectWarning(t *testing.T) {
	// Module empty + no detectable git user -> DetectModule placeholder + warning.
	t.Setenv("GITHUB_USER", "")
	t.Setenv("GH_USER", "")
	t.Setenv("HOME", t.TempDir())
	t.Setenv("GIT_CONFIG_GLOBAL", "/dev/null")
	t.Setenv("PATH", t.TempDir()) // hide git so gitUserName fails

	featPath := writeTempFeatures(t)
	target := filepath.Join(t.TempDir(), "myapp")
	var out, errOut bytes.Buffer
	err := Run(&out, &errOut, Options{
		ProjectID:    "Myapp",
		FeaturesPath: featPath,
		Dir:          target,
		// Module intentionally empty to trigger DetectModule.
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(errOut.String(), "warning") {
		t.Errorf("expected placeholder warning on errOut, got %q", errOut.String())
	}
	// Manifest should still be created.
	if _, err := os.Stat(filepath.Join(target, "specs", "manifest.yaml")); err != nil {
		t.Errorf("manifest not created: %v", err)
	}
}

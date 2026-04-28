//ff:func feature=cli-init type=test-helper control=sequence
//ff:what assertSkeletonManifest — manifest.yaml contains expected apiVersion/name/module

package cliinit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertSkeletonManifest(t *testing.T, target string) {
	t.Helper()
	manifest, err := os.ReadFile(filepath.Join(target, "specs/manifest.yaml"))
	if err != nil {
		t.Fatalf("read manifest: %v", err)
	}
	body := string(manifest)
	if !strings.Contains(body, "apiVersion: yongol/v1") {
		t.Errorf("manifest missing yongol/v1 apiVersion:\n%s", body)
	}
	if !strings.Contains(body, "name: myapp") {
		t.Errorf("manifest missing normalized name=myapp:\n%s", body)
	}
	if !strings.Contains(body, "module: github.com/test/myapp") {
		t.Errorf("manifest missing module override:\n%s", body)
	}
}

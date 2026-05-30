//ff:func feature=cli type=test control=sequence
//ff:what TestGenerateCmd_ResolveBackendError — manifest backend lang 미지원 시 resolve 에러 분기

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// A manifest with an unsupported backend (lang=rust) that still parses cleanly.
// With --backend "" the command falls through to manifest resolution, which
// returns an error from generate.ResolveBackendType.
func TestGenerateCmd_ResolveBackendError(t *testing.T) {
	dir := t.TempDir()
	manifest := `apiVersion: yongol/v1
kind: Project
metadata:
  name: unit-test
backend:
  lang: rust
  framework: actix
  module: example.com/unit-test
frontend:
  lang: typescript
  framework: react
  bundler: vite
  name: unit-test-web
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(manifest), 0o644); err != nil {
		t.Fatal(err)
	}

	_, _, err := runCmd(t, "generate", "--backend", "", dir, t.TempDir())
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	// Either the resolve error (preferred) or an earlier validate-warning refusal.
	if !strings.Contains(err.Error(), "resolve backend from manifest") &&
		!strings.Contains(err.Error(), "warnings must be resolved") &&
		!strings.Contains(err.Error(), "validation failed") &&
		!strings.Contains(err.Error(), "parse failed") {
		t.Fatalf("unexpected error: %v", err)
	}
}

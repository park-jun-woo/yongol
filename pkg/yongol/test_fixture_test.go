//ff:func feature=orchestrator type=test control=sequence
//ff:what PhaseT08 테스트 공용 헬퍼 — tempdir specs 구성 + zenflow fixture 절대경로 해석
package yongol

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// minimalManifest is the smallest manifest.yaml accepted by the manifest
// loader. Used by SSOT detection tests that need KindConfig to be populated
// but do not exercise manifest semantics themselves.
const minimalManifest = `apiVersion: yongol/v1
kind: Project
metadata:
  name: unit-test
backend:
  lang: go
  framework: gin
  module: example.com/unit-test
frontend:
  lang: typescript
  framework: react
  bundler: vite
  name: unit-test-web
`

// newTmpSpecsDir creates an empty temp dir that tests can populate with
// partial SSOT layouts. Callers MkdirAll / WriteFile individually so that each
// test can assert a minimal matrix of presence states.
func newTmpSpecsDir(t *testing.T) string {
	t.Helper()
	return t.TempDir()
}

// writeFile is a tiny wrapper that fails the test on write error; keeps
// individual test bodies focused on the matrix they care about.
func writeFile(t *testing.T, path, body string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}

// findZenflowSpecsAbs resolves the absolute path of the shared zenflow fixture
// (`dummys/zenflow/try-02/specs`) relative to *this* test file, so the tests
// work regardless of the caller's CWD and survive repo relocation.
func findZenflowSpecsAbs(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("runtime.Caller(0) failed")
	}
	// thisFile: <repo>/pkg/yongol/test_fixture_test.go
	pkgYongol := filepath.Dir(thisFile)
	repoRoot := filepath.Dir(filepath.Dir(pkgYongol))
	specs := filepath.Join(repoRoot, "dummys", "zenflow", "try-02", "specs")
	if _, err := os.Stat(specs); err != nil {
		t.Skipf("zenflow fixture unavailable at %s: %v", specs, err)
	}
	return specs
}

// hasKind reports whether detected contains a DetectedSSOT for k.
func hasKind(detected []DetectedSSOT, k SSOTKind) (DetectedSSOT, bool) {
	for _, d := range detected {
		if d.Kind == k {
			return d, true
		}
	}
	return DetectedSSOT{}, false
}

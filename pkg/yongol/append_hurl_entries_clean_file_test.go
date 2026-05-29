//ff:func feature=orchestrator type=test control=sequence
//ff:what appendHurlEntries — 정상 hurl 파일은 diag 없이 entries 로 등록
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

// TestAppendHurlEntriesCleanFile ensures that a well-formed hurl file produces
// no diagnostics and does not trip the gate. This is the regression counter-
// part to the propagation test above.
func TestAppendHurlEntriesCleanFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "clean.hurl")
	body := "GET {{host}}/ping\nHTTP 200\n"
	if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	fs := &Fullstack{HurlFiles: []string{path}}
	appendHurlEntries(fs)

	if len(fs.ParseDiagnostics) != 0 {
		t.Fatalf("expected 0 diagnostics on a clean hurl file, got %d: %+v",
			len(fs.ParseDiagnostics), fs.ParseDiagnostics)
	}
	if len(fs.HurlEntries) == 0 {
		t.Fatalf("expected at least 1 entry for a clean hurl file, got 0")
	}
}

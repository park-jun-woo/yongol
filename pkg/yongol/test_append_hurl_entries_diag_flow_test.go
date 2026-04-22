//ff:func feature=orchestrator type=test control=sequence
//ff:what appendHurlEntries 가 hurl 파서 diag 를 fs.ParseDiagnostics 로 전파하는지 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

// TestAppendHurlEntriesPropagatesDiagnostics ensures that when hurl.ParseFile
// returns diagnostics (e.g. a missing file), appendHurlEntries forwards them
// into fs.ParseDiagnostics rather than silently dropping them. This guards
// the validate-gate which triggers on len(fs.ParseDiagnostics) > 0.
func TestAppendHurlEntriesPropagatesDiagnostics(t *testing.T) {
	fs := &Fullstack{
		HurlFiles: []string{"/nonexistent/does-not-exist.hurl"},
	}

	appendHurlEntries(fs)

	if len(fs.ParseDiagnostics) == 0 {
		t.Fatalf("expected ParseDiagnostics to contain hurl parse error, got 0")
	}
	if len(fs.HurlEntries) != 0 {
		t.Fatalf("expected 0 entries for a missing file, got %d", len(fs.HurlEntries))
	}
}

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

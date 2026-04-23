//ff:func feature=orchestrator type=test control=sequence
//ff:what appendHurlEntries 가 hurl 파서 diag 를 fs.ParseDiagnostics 로 전파하는지 검증
package yongol

import (
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

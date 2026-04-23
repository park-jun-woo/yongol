//ff:func feature=orchestrator type=test control=sequence
//ff:what ParseFile — 존재하지 않는 파일은 File/Phase/Level/Message 모든 진단 필드 채워 반환

package sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestParseFile_ReadError_DiagnosticComplete(t *testing.T) {
	specs, diags := ParseFile("/definitely/does/not/exist.sql")
	if specs != nil {
		t.Fatalf("want nil specs, got %v", specs)
	}
	if len(diags) != 1 {
		t.Fatalf("want 1 diag, got %d: %v", len(diags), diags)
	}
	d := diags[0]
	// PhaseP01 이후 Diagnostic 필드 완결성 회귀 방어
	if d.Phase != diagnostic.PhaseParse {
		t.Errorf("Phase = %q, want %q", d.Phase, diagnostic.PhaseParse)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level = %q, want %q", d.Level, diagnostic.LevelError)
	}
	if d.Message == "" {
		t.Error("Message is empty")
	}
	if d.File == "" {
		t.Error("File is empty")
	}
}

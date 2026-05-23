//ff:func feature=validate type=test control=sequence topic=response-body-required
//ff:what buildO05Diagnostic — 진단 메시지 필드·레벨·내용 검증

package openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestBuildO05Diagnostic(t *testing.T) {
	lines := &oapiparser.LineIndex{
		Operations: map[string]int{
			"getUser": 42,
		},
	}

	diag := buildO05Diagnostic("404", "getUser", "/users/{id}", lines)

	if diag.File != "api/openapi.yaml" {
		t.Errorf("File = %q, want api/openapi.yaml", diag.File)
	}
	if diag.Line != 42 {
		t.Errorf("Line = %d, want 42", diag.Line)
	}
	if diag.Phase != diagnostic.PhaseValidate {
		t.Errorf("Phase = %v, want PhaseValidate", diag.Phase)
	}
	if diag.Level != diagnostic.LevelError {
		t.Errorf("Level = %v, want LevelError", diag.Level)
	}
	if !strings.Contains(diag.Message, "O-5") {
		t.Errorf("Message missing O-5: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "404") {
		t.Errorf("Message missing status: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "getUser") {
		t.Errorf("Message missing opID: %s", diag.Message)
	}

	// Test with nil lines (line should be 0)
	diag2 := buildO05Diagnostic("500", "unknownOp", "/path", nil)
	if diag2.Line != 0 {
		t.Errorf("nil lines: Line = %d, want 0", diag2.Line)
	}
}

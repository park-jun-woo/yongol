//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what xdn06TypeMismatchDiag — 진단 메시지 필드·레벨·내용·advice 검증

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXdn06TypeMismatchDiag(t *testing.T) {
	def := pmanifest.ClaimDef{
		Key:        "id",
		GoType:     "string",
		SourceLine: 22,
	}
	col := ddl.Column{RawType: "BIGINT"}
	diag := xdn06TypeMismatchDiag("UserID", "users", def, col)

	if diag.File != "manifest.yaml" {
		t.Errorf("File = %q, want manifest.yaml", diag.File)
	}
	if diag.Line != 22 {
		t.Errorf("Line = %d, want 22", diag.Line)
	}
	if diag.Phase != diagnostic.PhaseValidate {
		t.Errorf("Phase = %v, want PhaseValidate", diag.Phase)
	}
	if diag.Level != diagnostic.LevelError {
		t.Errorf("Level = %v, want LevelError", diag.Level)
	}
	if !strings.Contains(diag.Message, "XDN-06") {
		t.Errorf("Message missing XDN-06: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "UserID") {
		t.Errorf("Message missing field name: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, `"string"`) {
		t.Errorf("Message missing claim type: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "BIGINT") {
		t.Errorf("Message missing DDL raw type: %s", diag.Message)
	}
	// Advice should suggest int64 for BIGINT
	if !strings.Contains(diag.Advice, "int64") {
		t.Errorf("Advice missing suggested type int64: %s", diag.Advice)
	}
}

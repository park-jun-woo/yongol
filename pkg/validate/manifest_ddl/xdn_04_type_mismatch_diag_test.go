//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what xdn04TypeMismatchDiag — 진단 메시지 필드·레벨·내용 검증

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXdn04TypeMismatchDiag(t *testing.T) {
	def := pmanifest.ClaimDef{
		Key:        "id",
		SourceLine: 15,
	}
	diag := xdn04TypeMismatchDiag("UserID", "users", "string", "int64", def)

	if diag.File != "manifest.yaml" {
		t.Errorf("File = %q, want manifest.yaml", diag.File)
	}
	if diag.Line != 15 {
		t.Errorf("Line = %d, want 15", diag.Line)
	}
	if diag.Phase != diagnostic.PhaseValidate {
		t.Errorf("Phase = %v, want PhaseValidate", diag.Phase)
	}
	if diag.Level != diagnostic.LevelError {
		t.Errorf("Level = %v, want LevelError", diag.Level)
	}
	if !strings.Contains(diag.Message, "XDN-04") {
		t.Errorf("Message missing XDN-04: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "UserID") {
		t.Errorf("Message missing field name: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, `"string"`) {
		t.Errorf("Message missing claim type: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, `"int64"`) {
		t.Errorf("Message missing DDL type: %s", diag.Message)
	}
	if !strings.Contains(diag.Advice, "int64") {
		t.Errorf("Advice missing DDL type: %s", diag.Advice)
	}
}

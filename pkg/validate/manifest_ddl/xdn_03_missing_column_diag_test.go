//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what xdn03MissingColumnDiag — 진단 메시지 필드·레벨·내용 검증

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXdn03MissingColumnDiag(t *testing.T) {
	def := pmanifest.ClaimDef{
		Key:        "org_id",
		SourceLine: 42,
	}
	diag := xdn03MissingColumnDiag("OrgID", "users", def)

	if diag.File != "manifest.yaml" {
		t.Errorf("File = %q, want manifest.yaml", diag.File)
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
	if !strings.Contains(diag.Message, "XDN-03") {
		t.Errorf("Message missing XDN-03: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "OrgID") {
		t.Errorf("Message missing field name OrgID: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "org_id") {
		t.Errorf("Message missing claim key org_id: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "users") {
		t.Errorf("Message missing table name users: %s", diag.Message)
	}
	if !strings.Contains(diag.Advice, "users") {
		t.Errorf("Advice missing table name: %s", diag.Advice)
	}
}

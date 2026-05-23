//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what xdn05MissingTypeDiag — 진단 메시지 필드·레벨·내용 검증

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXdn05MissingTypeDiag(t *testing.T) {
	def := pmanifest.ClaimDef{
		Key:        "org_id",
		SourceLine: 18,
	}
	diag := xdn05MissingTypeDiag("OrgID", def)

	if diag.File != "manifest.yaml" {
		t.Errorf("File = %q, want manifest.yaml", diag.File)
	}
	if diag.Line != 18 {
		t.Errorf("Line = %d, want 18", diag.Line)
	}
	if diag.Phase != diagnostic.PhaseValidate {
		t.Errorf("Phase = %v, want PhaseValidate", diag.Phase)
	}
	if diag.Level != diagnostic.LevelError {
		t.Errorf("Level = %v, want LevelError", diag.Level)
	}
	if !strings.Contains(diag.Message, "XDN-05") {
		t.Errorf("Message missing XDN-05: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "OrgID") {
		t.Errorf("Message missing field name: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, "col:type") {
		t.Errorf("Message missing format hint: %s", diag.Message)
	}
	if !strings.Contains(diag.Advice, "org_id") {
		t.Errorf("Advice missing claim key: %s", diag.Advice)
	}
}

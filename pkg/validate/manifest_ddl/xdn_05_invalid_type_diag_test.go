//ff:func feature=validate type=test control=sequence topic=manifest-infra
//ff:what xdn05InvalidTypeDiag — 진단 메시지 필드·레벨·내용 검증

package manifest_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXdn05InvalidTypeDiag(t *testing.T) {
	def := pmanifest.ClaimDef{
		Key:        "data",
		GoType:     "float64",
		SourceLine: 20,
	}
	diag := xdn05InvalidTypeDiag("Data", def)

	if diag.File != "manifest.yaml" {
		t.Errorf("File = %q, want manifest.yaml", diag.File)
	}
	if diag.Line != 20 {
		t.Errorf("Line = %d, want 20", diag.Line)
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
	if !strings.Contains(diag.Message, "Data") {
		t.Errorf("Message missing field name: %s", diag.Message)
	}
	if !strings.Contains(diag.Message, `"float64"`) {
		t.Errorf("Message missing unknown type: %s", diag.Message)
	}
	if !strings.Contains(diag.Advice, "string, int64") {
		t.Errorf("Advice missing allowed types: %s", diag.Advice)
	}
}

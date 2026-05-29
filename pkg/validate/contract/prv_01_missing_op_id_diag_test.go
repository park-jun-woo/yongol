//ff:func feature=validate-contract type=test control=sequence
//ff:what TestPRV01MissingOpIDDiag — operationId 가 사라진 preserved 파일 Diagnostic 생성 검증

package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPRV01MissingOpIDDiag(t *testing.T) {
	d := prv01MissingOpIDDiag("svc/foo.go", "ActivateWorkflow", "activateWorkflow")
	if d.Level != diagnostic.LevelError || d.Phase != diagnostic.PhaseValidate {
		t.Errorf("level/phase = %v/%v", d.Level, d.Phase)
	}
	if !strings.Contains(d.Message, "[PRV-01]") {
		t.Errorf("message = %q", d.Message)
	}
	if !strings.Contains(d.Message, "ActivateWorkflow") || !strings.Contains(d.Message, "activateWorkflow") {
		t.Errorf("expected funcName and opID in message: %q", d.Message)
	}
	if d.Advice == "" {
		t.Error("expected non-empty advice")
	}
}

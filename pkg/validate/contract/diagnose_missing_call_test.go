//ff:func feature=validate-contract type=test control=sequence
//ff:what TestDiagnoseMissingCall — 사라진 @call/func 대상 참조 PRV-02 Diagnostic 생성 검증

package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestDiagnoseMissingCall(t *testing.T) {
	d := diagnoseMissingCall("svc/foo.go", "billing.CheckCredits")
	if d.Level != diagnostic.LevelError || d.Phase != diagnostic.PhaseValidate {
		t.Errorf("level/phase = %v/%v", d.Level, d.Phase)
	}
	if !strings.Contains(d.Message, "[PRV-02]") || !strings.Contains(d.Message, "billing.CheckCredits") {
		t.Errorf("message = %q", d.Message)
	}
	if d.Advice == "" {
		t.Error("expected non-empty advice")
	}
}

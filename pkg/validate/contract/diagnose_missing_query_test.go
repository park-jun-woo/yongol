//ff:func feature=validate-contract type=test control=sequence
//ff:what TestDiagnoseMissingQuery — 사라진 sqlc 쿼리 참조 PRV-02 Diagnostic 생성 검증

package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestDiagnoseMissingQuery(t *testing.T) {
	d := diagnoseMissingQuery("svc/foo.go", "UserDeleteByID")
	if d.Level != diagnostic.LevelError || d.Phase != diagnostic.PhaseValidate {
		t.Errorf("level/phase = %v/%v", d.Level, d.Phase)
	}
	if !strings.Contains(d.Message, "[PRV-02]") || !strings.Contains(d.Message, "UserDeleteByID") {
		t.Errorf("message = %q", d.Message)
	}
	if !strings.Contains(d.Advice, "UserDeleteByID") {
		t.Errorf("advice should reference query name: %q", d.Advice)
	}
}

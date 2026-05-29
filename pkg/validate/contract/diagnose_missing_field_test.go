//ff:func feature=validate-contract type=test control=sequence
//ff:what TestDiagnoseMissingField — 사라진 DDL 컬럼 필드 참조 PRV-02 Diagnostic 생성 검증

package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestDiagnoseMissingField(t *testing.T) {
	d := diagnoseMissingField("svc/foo.go", "u.DeletedAt")
	if d.File != "svc/foo.go" {
		t.Errorf("file = %q", d.File)
	}
	if d.Level != diagnostic.LevelError || d.Phase != diagnostic.PhaseValidate {
		t.Errorf("level/phase = %v/%v", d.Level, d.Phase)
	}
	if !strings.Contains(d.Message, "[PRV-02]") || !strings.Contains(d.Message, "DeletedAt") {
		t.Errorf("message = %q", d.Message)
	}
	if !strings.Contains(d.Message, "u.DeletedAt") {
		t.Errorf("expected full selector in message: %q", d.Message)
	}

	t.Run("no dot selector", func(t *testing.T) {
		d := diagnoseMissingField("x.go", "Email")
		if !strings.Contains(d.Message, "Email") {
			t.Errorf("message = %q", d.Message)
		}
	})
}

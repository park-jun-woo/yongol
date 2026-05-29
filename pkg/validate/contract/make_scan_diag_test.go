//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestMakeScanDiag — PRV-13 Diagnostic 값 생성(메시지/라인) 검증

package contract

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestMakeScanDiag(t *testing.T) {
	d := makeScanDiag("svc/foo.go", 42)
	if d.File != "svc/foo.go" || d.Line != 42 {
		t.Errorf("file/line = %q/%d", d.File, d.Line)
	}
	if d.Level != diagnostic.LevelError || d.Phase != diagnostic.PhaseValidate {
		t.Errorf("level/phase = %v/%v", d.Level, d.Phase)
	}
	if !strings.Contains(d.Message, "[PRV-13]") || !strings.Contains(d.Message, "42") {
		t.Errorf("message = %q", d.Message)
	}
	if !strings.Contains(d.Advice, "nolint:prv-13") {
		t.Errorf("advice = %q", d.Advice)
	}
}

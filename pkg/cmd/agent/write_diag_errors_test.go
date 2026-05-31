//ff:func feature=agent type=test control=sequence
//ff:what TestWriteDiagErrors — 진단 메시지/Advice를 문자열 빌더에 형식대로 기록 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestWriteDiagErrors(t *testing.T) {
	diags := []diagnostic.Diagnostic{
		{Message: "bad type", Advice: "use BIGINT"},
		{Message: "missing field"}, // no advice
	}
	var b strings.Builder
	writeDiagErrors(&b, diags)
	out := b.String()

	if !strings.HasPrefix(out, "Validate errors:\n") {
		t.Errorf("missing header: %q", out)
	}
	if !strings.Contains(out, "- bad type\n") || !strings.Contains(out, "  Advice: use BIGINT\n") {
		t.Errorf("missing message/advice: %q", out)
	}
	if !strings.Contains(out, "- missing field\n") {
		t.Errorf("missing second message: %q", out)
	}
	if strings.Count(out, "Advice:") != 1 {
		t.Errorf("advice should appear once (only when non-empty): %q", out)
	}
}

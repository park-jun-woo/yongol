//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestXfs39_NonBuiltinFallbackAdvice — non-builtin 패키지 @call 시 generic advice 확인

package ssac_func

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXfs39_NonBuiltinFallbackAdvice verifies that XFS-39 advice for a
// non-builtin package falls back to the generic "Define function under pkg/"
// message.
func TestXfs39_NonBuiltinFallbackAdvice(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "HandleOrder",
			FileName: "service/order/handle.ssac",
			Sequences: []parsessac.Sequence{{
				Type:  "call",
				Model: "billing.DoSomething",
				Line:  5,
			}},
		}},
	}
	fs.SetGround(ground.Build(fs))
	diags := xfs39CallToFuncSpec(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Advice, "Define function") {
		t.Errorf("expected generic advice, got %q", diags[0].Advice)
	}
}

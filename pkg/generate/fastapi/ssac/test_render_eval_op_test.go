//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderEvalOpNoPrefix — renderEvalOp 모듈 접두사 미사용 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderEvalOpNoPrefix(t *testing.T) {
	t.Run("NoModulePrefix", func(t *testing.T) {
		op := &ir.EvalOp{
			Package:    "billing",
			Function:   "IsZeroBalance",
			StatusCode: 402,
			Message:    "Insufficient balance",
			Args: []ir.FieldArg{
				{Source: "org", Field: ".CreditsBalance", Location: ir.LocVar},
			},
		}
		var b strings.Builder
		renderEvalOp(&b, op, "    ")
		got := b.String()
		// Must NOT have "billing." prefix
		if strings.Contains(got, "billing.") {
			t.Errorf("expected no module prefix, got: %s", got)
		}
		if !strings.Contains(got, "if is_zero_balance(") {
			t.Errorf("expected direct function call, got: %s", got)
		}
		if !strings.Contains(got, "status_code=402") {
			t.Errorf("expected status code, got: %s", got)
		}
	})

	t.Run("NilOp", func(t *testing.T) {
		var b strings.Builder
		renderEvalOp(&b, nil, "    ")
		if b.Len() != 0 {
			t.Error("expected empty for nil op")
		}
	})
}

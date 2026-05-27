//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderCallOpNoPrefix — renderCallOp 모듈 접두사 미사용 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderCallOpNoPrefix(t *testing.T) {
	t.Run("NoModulePrefix", func(t *testing.T) {
		op := &ir.CallOp{
			Package:   "billing",
			Function:  "IsZeroBalance",
			ResultVar: "result",
			Args: []ir.FieldArg{
				{Source: "org", Field: ".CreditsBalance", Location: ir.LocVar},
			},
		}
		var b strings.Builder
		renderCallOp(&b, op, "    ")
		got := b.String()
		// Must NOT have "billing." prefix
		if strings.Contains(got, "billing.") {
			t.Errorf("expected no module prefix, got: %s", got)
		}
		if !strings.Contains(got, "result = await is_zero_balance(") {
			t.Errorf("expected direct function call, got: %s", got)
		}
	})

	t.Run("NoResultVar", func(t *testing.T) {
		op := &ir.CallOp{
			Package:  "worker",
			Function: "ProcessAction",
		}
		var b strings.Builder
		renderCallOp(&b, op, "    ")
		got := b.String()
		if strings.Contains(got, "worker.") {
			t.Errorf("expected no module prefix, got: %s", got)
		}
		if !strings.Contains(got, "await process_action()") {
			t.Errorf("expected direct function call, got: %s", got)
		}
	})

	t.Run("NilOp", func(t *testing.T) {
		var b strings.Builder
		renderCallOp(&b, nil, "    ")
		if b.Len() != 0 {
			t.Error("expected empty for nil op")
		}
	})
}

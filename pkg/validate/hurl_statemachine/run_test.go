//ff:func feature=validate type=test control=sequence topic=hurl-statemachine
//ff:what Run — Hurl↔StateMachine 교차 검증 실행 통합 검증

package hurl_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	t.Run("nil_fs_no_panic", func(t *testing.T) {
		diags := Run(nil)
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})

	t.Run("empty_fs_no_diag", func(t *testing.T) {
		diags := Run(&yongol.Fullstack{})
		if len(diags) != 0 {
			t.Errorf("expected 0 diags, got %d", len(diags))
		}
	})
}

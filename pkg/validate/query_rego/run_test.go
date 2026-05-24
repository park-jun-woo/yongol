//ff:func feature=validate type=test control=sequence topic=query-rego
//ff:what Run — 전체 XQP-* 실행 (빈 fs) 검증

package query_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun(t *testing.T) {
	t.Run("empty fullstack returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}

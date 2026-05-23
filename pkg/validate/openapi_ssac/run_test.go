//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what Run — empty Fullstack 빈 결과 검증

package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRunOpenAPISsac(t *testing.T) {
	t.Run("empty fullstack returns empty", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}

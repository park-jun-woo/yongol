//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=response
//ff:what buildFuncFieldLookup 단위 테스트 (JSONName 직접 매칭 + PascalToSnake fallback)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestBuildFuncFieldLookup(t *testing.T) {
	t.Run("nil spec → empty map", func(t *testing.T) {
		got := buildFuncFieldLookup(nil)
		if got == nil || len(got) != 0 {
			t.Errorf("expected empty non-nil map, got %v", got)
		}
	})
	t.Run("explicit JSONName and snake fallback", func(t *testing.T) {
		spec := &funcspec.FuncSpec{
			ResponseFields: []funcspec.Field{
				{Name: "Summary", JSONName: "result_summary"},
				{Name: "CreatedAt"}, // no JSONName → created_at
			},
		}
		got := buildFuncFieldLookup(spec)
		if got["result_summary"] != "Summary" {
			t.Errorf("result_summary → %q, want Summary", got["result_summary"])
		}
		if got["created_at"] != "CreatedAt" {
			t.Errorf("created_at → %q, want CreatedAt", got["created_at"])
		}
		if len(got) != 2 {
			t.Errorf("expected 2 entries, got %d (%v)", len(got), got)
		}
	})
}

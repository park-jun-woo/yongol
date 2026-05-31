//ff:func feature=gen-gogin type=test control=sequence
//ff:what collectOwnerships 단위 테스트 (ParsedPolicies 전체 @ownership 평탄화)
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCollectOwnerships(t *testing.T) {
	t.Run("nil fullstack → nil", func(t *testing.T) {
		if got := collectOwnerships(nil); got != nil {
			t.Errorf("expected nil, got %v", got)
		}
	})
	t.Run("flattens across policies in order", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ParsedPolicies: []rego.Policy{
				{Ownerships: []rego.OwnershipMapping{{Resource: "a"}, {Resource: "b"}}},
				{Ownerships: []rego.OwnershipMapping{{Resource: "c"}}},
			},
		}
		got := collectOwnerships(fs)
		if len(got) != 3 {
			t.Fatalf("expected 3 mappings, got %d", len(got))
		}
		wantOrder := []string{"a", "b", "c"}
		for i, w := range wantOrder {
			if got[i].Resource != w {
				t.Errorf("got[%d].Resource = %q, want %q", i, got[i].Resource, w)
			}
		}
	})
}

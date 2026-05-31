//ff:func feature=validate type=test control=sequence topic=policy-check
//ff:what TestSSaCRegoHelpers — unit tests for the pure ssac_rego helper functions
package ssac_rego

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectSSaCAuthPairs(t *testing.T) {
	funcs := []ssac.ServiceFunc{{
		Sequences: []ssac.Sequence{
			authSeq("delete", "project", 5),
			{Type: "get"}, // non-auth ignored
		},
	}}
	pairs := collectSSaCAuthPairs(funcs)
	if !pairs[[2]string{"delete", "project"}] {
		t.Errorf("pairs = %v", pairs)
	}
	if len(pairs) != 1 {
		t.Errorf("expected only one pair, got %v", pairs)
	}
}

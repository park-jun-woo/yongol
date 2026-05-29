//ff:func feature=rule type=test control=sequence
//ff:what TestCoverageCheck_MissingLookupKey_Violation — Lookup 에 키 자체가 없으면 미사용으로 간주

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestCoverageCheck_MissingLookupKey_Violation(t *testing.T) {
	ground := &Ground{
		Lookup: map[string]StringSet{},
	}
	ctx := toulmin.NewContext()
	ctx.Set("ground", ground)
	ctx.Set("claim", "anything")

	spec := &CoverageCheckSpec{
		BaseSpec:  BaseSpec{Rule: "CC-3", Level: "ERROR"},
		LookupKey: "missing.key",
	}
	ok, _ := CoverageCheck(ctx, toulmin.Specs{spec})
	if !ok {
		t.Fatalf("CoverageCheck(missing lookup) ok = false; want true (treated as unused)")
	}
}

//ff:func feature=rule type=test control=sequence
//ff:what TestCoverageCheck_Used_NoViolation — Ground.Lookup 에 claim 이 있으면 위반 없음

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestCoverageCheck_Used_NoViolation(t *testing.T) {
	ground := &Ground{
		Lookup: map[string]StringSet{
			"openapi.used": {"listUsers": true},
		},
	}
	ctx := toulmin.NewContext()
	ctx.Set("ground", ground)
	ctx.Set("claim", "listUsers")

	spec := &CoverageCheckSpec{
		BaseSpec:  BaseSpec{Rule: "CC-1", Level: "WARNING", Message: "unused"},
		LookupKey: "openapi.used",
	}
	ok, ev := CoverageCheck(ctx, toulmin.Specs{spec})
	if ok || ev != nil {
		t.Fatalf("CoverageCheck(used) = (%v, %v); want (false, nil)", ok, ev)
	}
}

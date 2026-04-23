//ff:func feature=rule type=test control=sequence
//ff:what TestCoverageCheck_Unused_Violation — Ground.Lookup 에 claim 이 없으면 Evidence 방출

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestCoverageCheck_Unused_Violation(t *testing.T) {
	ground := &Ground{
		Lookup: map[string]StringSet{
			"openapi.used": {"listUsers": true},
		},
	}
	ctx := toulmin.NewContext()
	ctx.Set("ground", ground)
	ctx.Set("claim", "deleteUsers")

	spec := &CoverageCheckSpec{
		BaseSpec:  BaseSpec{Rule: "CC-2", Level: "WARNING", Message: "unused"},
		LookupKey: "openapi.used",
	}
	ok, ev := CoverageCheck(ctx, toulmin.Specs{spec})
	if !ok {
		t.Fatalf("CoverageCheck(unused) ok = false; want true")
	}
	e, okType := ev.(*Evidence)
	if !okType {
		t.Fatalf("CoverageCheck evidence type = %T; want *Evidence", ev)
	}
	if e.Rule != "CC-2" || e.Ref != "deleteUsers" {
		t.Fatalf("CoverageCheck evidence = %+v; want Rule=CC-2 Ref=deleteUsers", e)
	}
}

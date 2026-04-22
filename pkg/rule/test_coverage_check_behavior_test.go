//ff:func feature=rule type=test control=sequence
//ff:what CoverageCheck 실제 동작 경로 — Ground.Lookup[key] 기준으로 사용 여부 검증

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

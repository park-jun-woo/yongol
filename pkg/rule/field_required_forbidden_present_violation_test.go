//ff:func feature=rule type=test control=sequence
//ff:what TestFieldRequired_ForbiddenPresent_Violation — Required=false 인데 금지 필드가 있으면 Evidence 방출

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestFieldRequired_ForbiddenPresent_Violation(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set("claim", map[string]bool{"legacy": true})

	spec := &FieldRequiredSpec{
		BaseSpec: BaseSpec{Rule: "FR-3", Level: "WARNING"},
		Field:    "legacy",
		Required: false,
	}
	ok, ev := FieldRequired(ctx, toulmin.Specs{spec})
	if !ok {
		t.Fatalf("FieldRequired(forbidden+present) ok = false; want true")
	}
	e, okType := ev.(*Evidence)
	if !okType || e.Ref != "legacy" {
		t.Fatalf("FieldRequired evidence = %+v; want Ref=legacy", e)
	}
}

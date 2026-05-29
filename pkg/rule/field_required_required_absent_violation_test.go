//ff:func feature=rule type=test control=sequence
//ff:what TestFieldRequired_RequiredAbsent_Violation — Required=true 이고 필드가 없으면 Evidence 방출

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestFieldRequired_RequiredAbsent_Violation(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set("claim", map[string]bool{"other": true})

	spec := &FieldRequiredSpec{
		BaseSpec: BaseSpec{Rule: "FR-2", Level: "ERROR", Message: "missing"},
		Field:    "name",
		Required: true,
	}
	ok, ev := FieldRequired(ctx, toulmin.Specs{spec})
	if !ok {
		t.Fatalf("FieldRequired(required+absent) ok = false; want true")
	}
	e, okType := ev.(*Evidence)
	if !okType {
		t.Fatalf("FieldRequired evidence type = %T; want *Evidence", ev)
	}
	if e.Rule != "FR-2" || e.Ref != "name" {
		t.Fatalf("FieldRequired evidence = %+v; want Rule=FR-2 Ref=name", e)
	}
}

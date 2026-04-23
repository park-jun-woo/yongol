//ff:func feature=rule type=test control=sequence
//ff:what TestFieldRequired_ForbiddenAbsent_NoViolation — Required=false 이고 금지 필드가 없으면 위반 없음

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestFieldRequired_ForbiddenAbsent_NoViolation(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set("claim", map[string]bool{})

	spec := &FieldRequiredSpec{
		BaseSpec: BaseSpec{Rule: "FR-4", Level: "WARNING"},
		Field:    "legacy",
		Required: false,
	}
	ok, ev := FieldRequired(ctx, toulmin.Specs{spec})
	if ok || ev != nil {
		t.Fatalf("FieldRequired(forbidden+absent) = (%v, %v); want (false, nil)", ok, ev)
	}
}

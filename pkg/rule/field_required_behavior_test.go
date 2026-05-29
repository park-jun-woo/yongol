//ff:func feature=rule type=test control=sequence
//ff:what TestFieldRequired_RequiredPresent_NoViolation — Required=true 이고 필드가 있으면 위반 없음

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestFieldRequired_RequiredPresent_NoViolation(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set("claim", map[string]bool{"name": true})

	spec := &FieldRequiredSpec{
		BaseSpec: BaseSpec{Rule: "FR-1", Level: "ERROR"},
		Field:    "name",
		Required: true,
	}
	ok, ev := FieldRequired(ctx, toulmin.Specs{spec})
	if ok || ev != nil {
		t.Fatalf("FieldRequired(required+present) = (%v, %v); want (false, nil)", ok, ev)
	}
}

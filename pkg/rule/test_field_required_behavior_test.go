//ff:func feature=rule type=test control=sequence
//ff:what FieldRequired 실제 동작 경로 — Required=true/false 각각의 present/absent 케이스

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

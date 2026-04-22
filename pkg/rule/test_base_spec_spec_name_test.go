//ff:func feature=rule type=test control=sequence
//ff:what BaseSpec.SpecName — Rule 필드를 그대로 반환하는지 검증

package rule

import "testing"

func TestBaseSpec_SpecName(t *testing.T) {
	cases := []struct {
		in   BaseSpec
		want string
	}{
		{BaseSpec{Rule: "V11-1", Level: "ERROR"}, "V11-1"},
		{BaseSpec{Rule: "", Level: "ERROR"}, ""},
		{BaseSpec{Rule: "CC-usage", Level: "WARNING", Message: "m"}, "CC-usage"},
	}
	for _, c := range cases {
		if got := c.in.SpecName(); got != c.want {
			t.Errorf("SpecName(%+v) = %q; want %q", c.in, got, c.want)
		}
	}
}

func TestBaseSpec_SpecName_EmbeddedInConcrete(t *testing.T) {
	vd := VarDeclaredSpec{BaseSpec: BaseSpec{Rule: "VD-1", Level: "ERROR"}}
	if vd.SpecName() != "VD-1" {
		t.Fatalf("VarDeclaredSpec.SpecName() = %q; want VD-1", vd.SpecName())
	}
	cc := CoverageCheckSpec{BaseSpec: BaseSpec{Rule: "CC-1", Level: "WARNING"}, LookupKey: "k"}
	if cc.SpecName() != "CC-1" {
		t.Fatalf("CoverageCheckSpec.SpecName() = %q; want CC-1", cc.SpecName())
	}
	fr := FieldRequiredSpec{BaseSpec: BaseSpec{Rule: "FR-1", Level: "ERROR"}, Field: "x", Required: true}
	if fr.SpecName() != "FR-1" {
		t.Fatalf("FieldRequiredSpec.SpecName() = %q; want FR-1", fr.SpecName())
	}
}

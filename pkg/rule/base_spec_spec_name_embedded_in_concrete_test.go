//ff:func feature=rule type=test control=sequence
//ff:what TestBaseSpec_SpecName_EmbeddedInConcrete — BaseSpec 을 임베드한 VarDeclared/Coverage/FieldRequiredSpec 이 SpecName 을 승계하는지

package rule

import "testing"

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

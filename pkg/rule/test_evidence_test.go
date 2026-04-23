//ff:func feature=rule type=test control=sequence
//ff:what TestEvidence_FieldsAssignable — Evidence 구조체 필드에 할당/비교가 기대대로 동작

package rule

import "testing"

func TestEvidence_FieldsAssignable(t *testing.T) {
	e := Evidence{
		Rule:    "R-1",
		Level:   "ERROR",
		Ref:     "field.name",
		Message: "missing field",
	}
	if e.Rule != "R-1" || e.Level != "ERROR" || e.Ref != "field.name" || e.Message != "missing field" {
		t.Fatalf("Evidence fields mismatch: %+v", e)
	}
}

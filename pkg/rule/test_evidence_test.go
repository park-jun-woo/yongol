//ff:func feature=rule type=test control=sequence
//ff:what Evidence — 규칙 위반 결과 구조체 필드 할당/비교 기본 동작 검증

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

func TestEvidence_ZeroValue(t *testing.T) {
	var e Evidence
	if e.Rule != "" || e.Level != "" || e.Ref != "" || e.Message != "" {
		t.Fatalf("zero Evidence not empty: %+v", e)
	}
}

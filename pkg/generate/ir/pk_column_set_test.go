//ff:func feature=gen-ir type=test control=sequence
//ff:what TestPkColumnSet -- PK 컬럼 슬라이스→집합 변환 및 멤버십 검증

package ir

import "testing"

func TestPkColumnSet(t *testing.T) {
	set := pkColumnSet([]string{"id", "tenant_id"})
	if len(set) != 2 {
		t.Fatalf("len(set) = %d, want 2", len(set))
	}
	if !set["id"] || !set["tenant_id"] {
		t.Errorf("set = %v, want id and tenant_id present", set)
	}
	if set["missing"] {
		t.Error("set should not contain missing key")
	}

	// empty input -> empty set
	empty := pkColumnSet(nil)
	if len(empty) != 0 {
		t.Errorf("len(empty) = %d, want 0", len(empty))
	}
}

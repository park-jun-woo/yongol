//ff:func feature=gen-ir type=test control=sequence
//ff:what ir 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ir

import (
	"testing"
)

func TestEnrichFieldArgLocations_ZeroCov(t *testing.T) {
	ops := []Op{
		{Kind: OpGet, Get: &GetOp{Args: []FieldArg{
			{Key: "id", Source: "request", Field: ".ID"},
			{Source: "currentUser", Field: ".OrgID"},
		}}},
	}
	enrichFieldArgLocations(ops, map[string]bool{".ID": true}, map[string]bool{})
	if ops[0].Get.Args[0].Location != LocPath {
		t.Errorf("expected LocPath, got %q", ops[0].Get.Args[0].Location)
	}
	if ops[0].Get.Args[1].Location != LocUser {
		t.Errorf("expected LocUser, got %q", ops[0].Get.Args[1].Location)
	}
}

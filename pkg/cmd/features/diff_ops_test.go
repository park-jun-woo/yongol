//ff:func feature=features type=test control=sequence
//ff:what TestDiffOps — op 집합 diff: 신규(Added)·삭제(Removed)·교집합 무변동 분기 검증
package features

import (
	"testing"

	featparser "github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestDiffOps(t *testing.T) {
	old := []featparser.Feature{
		{Op: "A"},
		{Op: "B"},
	}
	nw := []featparser.Feature{
		{Op: "B"}, // unchanged
		{Op: "C"}, // added
	}
	res := DiffOps(old, nw)

	if len(res.Added) != 1 || res.Added[0].Op != "C" {
		t.Errorf("Added = %v, want [C]", res.Added)
	}
	if len(res.Removed) != 1 || res.Removed[0].Op != "A" {
		t.Errorf("Removed = %v, want [A]", res.Removed)
	}
}

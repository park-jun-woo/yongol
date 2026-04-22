//ff:func feature=rule type=test control=sequence
//ff:what FieldRequired가 빈 specs에서 panic 없이 (false, nil)을 반환하는지 검증

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestFieldRequired_EmptySpecs(t *testing.T) {
	ok, ev := FieldRequired(toulmin.NewContext(), toulmin.Specs{})
	if ok || ev != nil {
		t.Fatalf("FieldRequired(empty) = (%v, %v); want (false, nil)", ok, ev)
	}
}

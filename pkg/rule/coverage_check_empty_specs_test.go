//ff:func feature=rule type=test control=sequence
//ff:what CoverageCheck가 빈 specs에서 panic 없이 (false, nil)을 반환하는지 검증

package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/pkg/toulmin"
)

func TestCoverageCheck_EmptySpecs(t *testing.T) {
	ok, ev := CoverageCheck(toulmin.NewContext(), toulmin.Specs{})
	if ok || ev != nil {
		t.Fatalf("CoverageCheck(empty) = (%v, %v); want (false, nil)", ok, ev)
	}
}

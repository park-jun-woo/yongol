//ff:func feature=features type=test control=sequence
//ff:what TestDiffOps — op 집합 diff: 신규(Added)·삭제(Removed)·교집합 무변동 분기 검증
package features

import (
	"testing"
)

func TestDiffOps_Empty(t *testing.T) {
	res := DiffOps(nil, nil)
	if len(res.Added) != 0 || len(res.Removed) != 0 {
		t.Errorf("empty diff produced %+v", res)
	}
}

//ff:func feature=contract type=test control=sequence
//ff:what TestStripAnnotationBlockBranches — 빈입력/주석없음/마지막줄 개행없는 주석 분기 검증
package contract

import (
	"testing"
)

func TestStripAnnotationBlock_Empty(t *testing.T) {
	if got := StripAnnotationBlock([]byte{}); len(got) != 0 {
		t.Errorf("empty input -> %q, want empty", got)
	}
}

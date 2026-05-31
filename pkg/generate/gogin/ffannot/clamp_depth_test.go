//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestClampDepth — 음수는 0, 비음수는 그대로 반환 검증
package ffannot

import (
	"testing"
)

func TestClampDepth(t *testing.T) {
	cases := map[int]int{-5: 0, -1: 0, 0: 0, 1: 1, 7: 7}
	for in, want := range cases {
		if got := clampDepth(in); got != want {
			t.Errorf("clampDepth(%d) = %d, want %d", in, got, want)
		}
	}
}

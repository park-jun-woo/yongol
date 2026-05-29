//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what methodGen.resolveErrCode 단위 테스트 (요청 status를 그대로 반환)

package ssac

import "testing"

func TestMethodGenResolveErrCode(t *testing.T) {
	g := &methodGen{}
	for _, want := range []int{0, 404, 409, 500} {
		if got := g.resolveErrCode(want); got != want {
			t.Errorf("resolveErrCode(%d) = %d, want %d", want, got, want)
		}
	}
}

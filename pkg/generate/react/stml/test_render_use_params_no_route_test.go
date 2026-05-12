//ff:func feature=stml-gen type=test control=sequence
//ff:what renderUseParamsWithRoute 라우트/파라미터 모두 없을 때 빈 문자열 검증

package stml

import "testing"

func TestRenderUseParamsWithRoute_NoRouteNoParams(t *testing.T) {
	got := renderUseParamsWithRoute(nil, "")
	if got != "" {
		t.Errorf("got %q, want empty", got)
	}
}

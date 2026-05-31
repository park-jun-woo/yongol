//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"
)

func TestRoutesFor_ZeroCov(t *testing.T) {
	if routesFor(bnFS(nil, nil)) != nil {
		t.Error("placeholder should be nil")
	}
}

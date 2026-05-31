//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestMiddlewaresFor_ZeroCov(t *testing.T) {
	if middlewaresFor(nil) != nil {
		t.Error("nil should be nil")
	}
	mc := &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Middleware: []string{"cors", "auth"}}}
	mws := middlewaresFor(bnFS(mc, nil))
	if len(mws) != 2 || mws[0].Name != "cors" {
		t.Errorf("middlewares wrong: %#v", mws)
	}
}

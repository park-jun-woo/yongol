//ff:func feature=gen-nestjs type=test control=iteration dimension=1
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"
)

func TestNestHTTPDecorator_ZeroCov(t *testing.T) {
	for m, want := range map[string]string{"get": "Get", "POST": "Post", "put": "Put", "delete": "Delete", "patch": "Patch", "weird": "Get"} {
		if got := nestHTTPDecorator(m); got != want {
			t.Errorf("nestHTTPDecorator(%q)=%q want %q", m, got, want)
		}
	}
}

//ff:func feature=external type=test control=sequence
//ff:what external 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package external

import (
	"testing"
)

func TestExtractPathMethods_ZeroCov(t *testing.T) {
	doc := sampleDoc()
	pi := doc.Paths.Map()["/items/{item_id}"]
	methods := extractPathMethods(pi, "/items/{item_id}")
	if len(methods) != 1 || methods[0].HTTPMethod != "GET" {
		t.Errorf("expected one GET method, got %#v", methods)
	}
}

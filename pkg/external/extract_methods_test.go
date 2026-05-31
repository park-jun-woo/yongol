//ff:func feature=external type=test control=sequence
//ff:what TestExtractMethods/buildMethodInfo/extractResponseTypes — 메서드·응답타입 추출 검증
package external

import (
	"testing"
)

func TestExtractMethods(t *testing.T) {
	doc := sampleDoc()
	methods := extractMethods(doc)
	// sorted by path: /items (POST create_item), /items/{item_id} (GET get_item)
	if len(methods) != 2 {
		t.Fatalf("expected 2 methods, got %d: %+v", len(methods), methods)
	}
	if methods[0].Name != "CreateItem" || methods[0].HTTPMethod != "POST" {
		t.Errorf("methods[0] = %+v, want CreateItem/POST", methods[0])
	}
	if methods[1].Name != "GetItem" || methods[1].HTTPMethod != "GET" {
		t.Errorf("methods[1] = %+v, want GetItem/GET", methods[1])
	}
}

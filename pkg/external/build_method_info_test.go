//ff:func feature=external type=test control=sequence
//ff:what TestExtractMethods/buildMethodInfo/extractResponseTypes — 메서드·응답타입 추출 검증
package external

import (
	"testing"
)

func TestBuildMethodInfo(t *testing.T) {
	doc := sampleDoc()
	getOp := doc.Paths.Map()["/items/{item_id}"].Get

	m := buildMethodInfo(getOp, "GET", "/items/{item_id}")
	if m.Name != "GetItem" {
		t.Errorf("Name = %q, want GetItem", m.Name)
	}
	if m.HTTPMethod != "GET" || m.Path != "/items/{item_id}" {
		t.Errorf("method/path mismatch: %+v", m)
	}
	if m.ReturnType != "GetItemResponse" {
		t.Errorf("ReturnType = %q, want GetItemResponse", m.ReturnType)
	}
	if len(m.Params) != 1 || m.Params[0].Name != "itemID" {
		t.Errorf("params = %+v", m.Params)
	}
}

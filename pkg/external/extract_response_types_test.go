//ff:func feature=external type=test control=sequence
//ff:what TestExtractMethods/buildMethodInfo/extractResponseTypes — 메서드·응답타입 추출 검증
package external

import (
	"testing"
)

func TestExtractResponseTypes(t *testing.T) {
	doc := sampleDoc()
	methods := extractMethods(doc)
	types := extractResponseTypes("svc", methods, doc)

	// Only get_item has a 200 JSON object response with properties.
	if len(types) != 1 {
		t.Fatalf("expected 1 response type, got %d: %+v", len(types), types)
	}
	st := types[0]
	if st.Name != "GetItemResponse" {
		t.Errorf("Name = %q, want GetItemResponse", st.Name)
	}
	if len(st.Fields) != 2 {
		t.Fatalf("expected 2 fields, got %+v", st.Fields)
	}
	// sortedKeys => id, name
	if st.Fields[0].Name != "ID" || st.Fields[0].JSONName != "id" || st.Fields[0].GoType != "int64" {
		t.Errorf("fields[0] = %+v", st.Fields[0])
	}
	if st.Fields[1].Name != "Name" || st.Fields[1].JSONName != "name" {
		t.Errorf("fields[1] = %+v", st.Fields[1])
	}
}

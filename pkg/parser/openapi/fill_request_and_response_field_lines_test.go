//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"testing"
)

func TestFillRequestAndResponseFieldLines(t *testing.T) {
	idx := newIdx()
	idx.RequestFields["op"] = map[string]int{"email": 42}
	idx.ResponseFields["op"] = map[string]int{"token": 99}

	req := map[string]FieldConstraint{"email": {Type: "string"}}
	fillRequestFieldLines(req, "op", idx)
	if req["email"].Line != 42 {
		t.Errorf("request line = %d, want 42", req["email"].Line)
	}

	resp := map[string]FieldConstraint{"token": {Type: "string"}}
	fillResponseFieldLines(resp, "op", idx)
	if resp["token"].Line != 99 {
		t.Errorf("response line = %d, want 99", resp["token"].Line)
	}
}

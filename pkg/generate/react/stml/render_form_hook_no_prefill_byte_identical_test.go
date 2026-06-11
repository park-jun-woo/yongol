//ff:func feature=stml-gen type=test control=sequence
//ff:what renderFormHook (d) prefill 부재 — respFields 무시·기존 출력 바이트 동일 (회귀 기준)

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// (d) no prefill → byte-identical to the pre-Phase035 emission (both branches).
func TestRenderFormHookNoPrefillByteIdentical(t *testing.T) {
	zodAction := stmlparser.ActionBlock{
		OperationID: "UpdateRule",
		Fields:      []stmlparser.FieldBind{{Name: "sheet_name"}, {Name: "start_row"}, {Name: "note"}},
	}
	// respFields passed but unused because Prefill == "".
	resp := map[string]oapiparser.FieldTypeInfo{"sheet_name": {Type: "string"}}
	withResp := renderFormHook(zodAction, prefillConstraints(), resp)
	without := renderFormHook(zodAction, prefillConstraints(), nil)
	if withResp != without {
		t.Fatalf("no-prefill zod output must ignore respFields:\n%q\nvs\n%q", withResp, without)
	}
	assertNotContains(t, withResp, "values:")
	assertNotContains(t, withResp, "resetOptions")

	untyped := stmlparser.ActionBlock{OperationID: "DeleteRoom", Fields: []stmlparser.FieldBind{{Name: "reason"}}}
	if got := renderFormHook(untyped, nil, resp); got != "const deleteRoomForm = useForm()" {
		t.Errorf("no-prefill untyped output = %q, want plain useForm()", got)
	}
}

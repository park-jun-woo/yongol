//ff:func feature=stml-gen type=test control=sequence
//ff:what renderFormHook (b) prefill+zod 일부 필드 응답 부재 — 빈 리터럴·나머지 data 참조(완전 객체 유지)

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// (b) prefill + zod, a field is absent from the response → empty literal, the
// rest stay data-referenced (complete object satisfies the exact `values` type).
func TestRenderFormHookPrefillPartialResponse(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "UpdateRule",
		Prefill:     "GetRule",
		Fields:      []stmlparser.FieldBind{{Name: "sheet_name"}, {Name: "start_row"}, {Name: "note"}},
	}
	// "note" is NOT in the response — must NOT reference getRuleData.note.
	resp := map[string]oapiparser.FieldTypeInfo{
		"sheet_name": {Type: "string"}, "start_row": {Type: "integer"},
	}
	code := renderFormHook(a, prefillConstraints(), resp)
	assertContains(t, code, "sheet_name: getRuleData.sheet_name ?? '',")
	assertContains(t, code, "start_row: getRuleData.start_row ?? 0,")
	assertContains(t, code, "note: '',")
	assertNotContains(t, code, "getRuleData.note")
}

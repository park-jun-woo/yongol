//ff:func feature=stml-gen type=test control=sequence
//ff:what renderFormHook (a) prefill+zod 전 필드 응답 존재 — 전부 data 참조 + 타입별 coalescing

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// (a) prefill + zod, all fields present in response → all data-referenced.
func TestRenderFormHookPrefillAllPresent(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "UpdateRule",
		Prefill:     "GetRule",
		Fields:      []stmlparser.FieldBind{{Name: "sheet_name"}, {Name: "start_row"}, {Name: "note"}},
	}
	resp := map[string]oapiparser.FieldTypeInfo{
		"sheet_name": {Type: "string"}, "start_row": {Type: "integer"}, "note": {Type: "string"},
	}
	code := renderFormHook(a, prefillConstraints(), resp)
	assertContains(t, code, "values: getRuleData")
	assertContains(t, code, "resetOptions: { keepDirtyValues: true }")
	assertContains(t, code, "sheet_name: getRuleData.sheet_name ?? '',")
	assertContains(t, code, "start_row: getRuleData.start_row ?? 0,")
	assertContains(t, code, "note: getRuleData.note ?? '',")
}

//ff:func feature=stml-gen type=test control=sequence
//ff:what renderFormHook (c) prefill+zod 무 — useForm({ values })에 응답 교집합만 partial 매핑

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// (c) prefill + no zod constraints → untyped useForm({ values }) with only the
// response-intersecting fields mapped (partial allowed for Record<string, any>).
func TestRenderFormHookPrefillUntyped(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "UpdateRule",
		Prefill:     "GetRule",
		Fields:      []stmlparser.FieldBind{{Name: "sheet_name"}, {Name: "note"}},
	}
	resp := map[string]oapiparser.FieldTypeInfo{"sheet_name": {Type: "string"}}
	code := renderFormHook(a, nil, resp)
	assertContains(t, code, "const updateRuleForm = useForm({")
	assertContains(t, code, "values: getRuleData")
	assertContains(t, code, "sheet_name: getRuleData.sheet_name,")
	assertContains(t, code, "resetOptions: { keepDirtyValues: true }")
	assertNotContains(t, code, "zodResolver")
	assertNotContains(t, code, "getRuleData.note")
}

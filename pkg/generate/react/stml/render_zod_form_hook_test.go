//ff:func feature=stml-gen type=test control=sequence
//ff:what renderZodFormHook — prefill 유/무 분기(values+resetOptions vs 기존 zodResolver) 검증

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderZodFormHook(t *testing.T) {
	fields := map[string]oapiparser.FieldConstraint{"sheet_name": {Type: "string", Required: true}}
	resp := map[string]oapiparser.FieldTypeInfo{"sheet_name": {Type: "string"}}

	withPrefill := stmlparser.ActionBlock{OperationID: "UpdateRule", Prefill: "GetRule"}
	code := renderZodFormHook(withPrefill, "updateRuleForm", fields, resp)
	assertContains(t, code, "const updateRuleSchema = z.object(")
	assertContains(t, code, "resolver: zodResolver(updateRuleSchema)")
	assertContains(t, code, "values: getRuleData")
	assertContains(t, code, "sheet_name: getRuleData.sheet_name ?? '',")
	assertContains(t, code, "resetOptions: { keepDirtyValues: true }")

	noPrefill := stmlparser.ActionBlock{OperationID: "UpdateRule"}
	plain := renderZodFormHook(noPrefill, "updateRuleForm", fields, resp)
	assertNotContains(t, plain, "values:")
	assertNotContains(t, plain, "resetOptions")
	assertContains(t, plain, "resolver: zodResolver(updateRuleSchema)")
}

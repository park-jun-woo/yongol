//ff:func feature=stml-gen type=test control=sequence
//ff:what renderUntypedFormHook — prefill 유/무, 교집합 없을 때 plain useForm() 폴백 검증

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUntypedFormHook(t *testing.T) {
	a := stmlparser.ActionBlock{OperationID: "UpdateRule", Prefill: "GetRule", Fields: []stmlparser.FieldBind{{Name: "sheet_name"}}}
	resp := map[string]oapiparser.FieldTypeInfo{"sheet_name": {Type: "string"}}
	code := renderUntypedFormHook(a, "updateRuleForm", resp)
	assertContains(t, code, "const updateRuleForm = useForm({")
	assertContains(t, code, "values: getRuleData")
	assertContains(t, code, "sheet_name: getRuleData.sheet_name,")

	// No prefill → plain useForm().
	noPrefill := stmlparser.ActionBlock{OperationID: "UpdateRule", Fields: a.Fields}
	if got := renderUntypedFormHook(noPrefill, "updateRuleForm", resp); got != "const updateRuleForm = useForm()" {
		t.Errorf("no prefill = %q", got)
	}

	// Prefill but no field overlap → plain useForm().
	if got := renderUntypedFormHook(a, "updateRuleForm", map[string]oapiparser.FieldTypeInfo{"other": {}}); got != "const updateRuleForm = useForm()" {
		t.Errorf("no overlap = %q", got)
	}
}

//ff:func feature=stml-gen type=test control=sequence
//ff:what renderFormHook 매칭되지 않는 operationId 시 plain useForm 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderFormHookNoMatchingOp(t *testing.T) {
	constraints := map[string]map[string]oapiparser.FieldConstraint{
		"OtherOp": {"field": {Type: "string", Required: true}},
	}
	a := stmlparser.ActionBlock{
		OperationID: "CreateWorkflow",
		Fields:      []stmlparser.FieldBind{{Name: "title"}},
	}
	code := renderFormHook(a, constraints)
	assertContains(t, code, "const createWorkflowForm = useForm()")
	assertNotContains(t, code, "zodResolver")
}

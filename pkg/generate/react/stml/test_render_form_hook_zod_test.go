//ff:func feature=stml-gen type=test control=sequence
//ff:what renderFormHook — zod 제약조건이 있을 때 zodResolver 적용을 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderFormHookWithZod(t *testing.T) {
	maxLen := 200
	constraints := map[string]map[string]oapiparser.FieldConstraint{
		"CreateWorkflow": {
			"title":         {Type: "string", Required: true, MaxLength: &maxLen},
			"trigger_event": {Type: "string", Required: true},
		},
	}
	a := stmlparser.ActionBlock{
		OperationID: "CreateWorkflow",
		Fields: []stmlparser.FieldBind{
			{Name: "title"},
			{Name: "trigger_event"},
		},
	}
	code := renderFormHook(a, constraints)
	assertContains(t, code, "const createWorkflowSchema = z.object(")
	assertContains(t, code, "resolver: zodResolver(createWorkflowSchema)")
	assertContains(t, code, "const createWorkflowForm = useForm(")
}

func TestRenderFormHookWithoutConstraints(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "DeleteRoom",
		Fields: []stmlparser.FieldBind{
			{Name: "reason"},
		},
	}
	code := renderFormHook(a, nil)
	assertContains(t, code, "const deleteRoomForm = useForm()")
	assertNotContains(t, code, "zodResolver")
	assertNotContains(t, code, "z.object")
}

func TestRenderFormHookNoMatchingOp(t *testing.T) {
	constraints := map[string]map[string]oapiparser.FieldConstraint{
		"OtherOp": {
			"field": {Type: "string", Required: true},
		},
	}
	a := stmlparser.ActionBlock{
		OperationID: "CreateWorkflow",
		Fields: []stmlparser.FieldBind{
			{Name: "title"},
		},
	}
	code := renderFormHook(a, constraints)
	assertContains(t, code, "const createWorkflowForm = useForm()")
	assertNotContains(t, code, "zodResolver")
}

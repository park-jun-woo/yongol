//ff:func feature=stml-gen type=test control=sequence
//ff:what renderFormHook zod 제약조건 시 zodResolver 적용 검증
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
		Fields:      []stmlparser.FieldBind{{Name: "title"}, {Name: "trigger_event"}},
	}
	code := renderFormHook(a, constraints)
	assertContains(t, code, "const createWorkflowSchema = z.object(")
	assertContains(t, code, "resolver: zodResolver(createWorkflowSchema)")
	assertContains(t, code, "const createWorkflowForm = useForm<z.infer<typeof createWorkflowSchema>>(")
}

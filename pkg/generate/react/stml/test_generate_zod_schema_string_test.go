//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema 문자열 필드 스키마 생성 검증
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestGenerateZodSchemaString(t *testing.T) {
	maxLen := 200
	fields := map[string]oapiparser.FieldConstraint{
		"title": {
			Type:      "string",
			Required:  true,
			MaxLength: &maxLen,
		},
		"trigger_event": {
			Type:     "string",
			Required: true,
		},
	}

	code := generateZodSchema("CreateWorkflow", fields)
	assertContains(t, code, "const createWorkflowSchema = z.object(")
	assertContains(t, code, "title: z.string().min(1).max(200)")
	assertContains(t, code, "trigger_event: z.string().min(1)")
}

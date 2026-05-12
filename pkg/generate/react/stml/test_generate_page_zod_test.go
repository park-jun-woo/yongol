//ff:func feature=stml-gen type=test control=sequence
//ff:what 폼 페이지 생성 시 zod 스키마 + zodResolver + import 통합 검증
package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePageWithZodSchema(t *testing.T) {
	page, _ := stmlparser.ParseReader("create-workflow-page.html", strings.NewReader(`<main>
  <div data-action="CreateWorkflow">
    <input data-field="title" placeholder="제목" />
    <input data-field="trigger_event" placeholder="트리거 이벤트" />
    <button type="submit">생성</button>
  </div>
</main>`))

	maxLen := 200
	opts := GenerateOptions{
		APIImportPath: "@/lib/api",
		RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
			"CreateWorkflow": {
				"title":         {Type: "string", Required: true, MaxLength: &maxLen},
				"trigger_event": {Type: "string", Required: true},
			},
		},
	}
	code := GeneratePage(page, "", opts)

	// Imports
	assertContains(t, code, "import { z } from 'zod'")
	assertContains(t, code, "import { zodResolver } from '@hookform/resolvers/zod'")
	assertContains(t, code, "import { useForm } from 'react-hook-form'")

	// Zod schema
	assertContains(t, code, "const createWorkflowSchema = z.object(")
	assertContains(t, code, "title: z.string().min(1).max(200)")
	assertContains(t, code, "trigger_event: z.string().min(1)")

	// zodResolver in useForm
	assertContains(t, code, "resolver: zodResolver(createWorkflowSchema)")
}

func TestGeneratePageWithoutConstraints(t *testing.T) {
	page, _ := stmlparser.ParseReader("edit-page.html", strings.NewReader(`<main>
  <div data-action="UpdateItem">
    <input data-field="Name" placeholder="이름" />
    <button type="submit">수정</button>
  </div>
</main>`))

	// No RequestConstraints → no zod
	code := GeneratePage(page, "")
	assertNotContains(t, code, "import { z }")
	assertNotContains(t, code, "zodResolver")
	assertNotContains(t, code, "z.object")
	assertContains(t, code, "const updateItemForm = useForm()")
}

func TestGeneratePageNoFieldsNoZod(t *testing.T) {
	page, _ := stmlparser.ParseReader("activate-page.html", strings.NewReader(`<main>
  <button data-action="ActivateWorkflow" data-param-workflow-id="route.WorkflowID">활성화</button>
</main>`))

	opts := GenerateOptions{
		APIImportPath: "@/lib/api",
		RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
			"ActivateWorkflow": {}, // empty body
		},
	}
	code := GeneratePage(page, "", opts)

	// No form → no zod
	assertNotContains(t, code, "import { z }")
	assertNotContains(t, code, "zodResolver")
	assertNotContains(t, code, "useForm")
}

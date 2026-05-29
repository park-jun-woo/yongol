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

	assertContains(t, code, "import { z } from 'zod'")
	assertContains(t, code, "import { zodResolver } from '@hookform/resolvers/zod'")
	assertContains(t, code, "import { useForm } from 'react-hook-form'")
	assertContains(t, code, "const createWorkflowSchema = z.object(")
	assertContains(t, code, "title: z.string().min(1).max(200)")
	assertContains(t, code, "trigger_event: z.string().min(1)")
	assertContains(t, code, "resolver: zodResolver(createWorkflowSchema)")
}

//ff:func feature=stml-gen type=test control=sequence
//ff:what 폼 필드 없는 액션 페이지에서 zod 미포함 검증
package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePageNoFieldsNoZod(t *testing.T) {
	page, _ := stmlparser.ParseReader("activate-page.html", strings.NewReader(`<main>
  <button data-action="ActivateWorkflow" data-param-workflow-id="route.WorkflowID">활성화</button>
</main>`))

	opts := GenerateOptions{
		APIImportPath: "@/lib/api",
		RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
			"ActivateWorkflow": {},
		},
	}
	code := GeneratePage(page, "", opts)

	assertNotContains(t, code, "import { z }")
	assertNotContains(t, code, "zodResolver")
	assertNotContains(t, code, "useForm")
}

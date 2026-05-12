//ff:func feature=stml-gen type=test control=sequence
//ff:what 액션 페이지에서 Button UI 컴포넌트 import 생성을 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestUIButtonImport(t *testing.T) {
	page, _ := stmlparser.ParseReader("action-page.html", strings.NewReader(`<main>
  <button data-action="ActivateWorkflow">활성화</button>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, "import { Button } from '@/components/ui/Button'")
	assertContains(t, code, "<Button ")
	assertNotContains(t, code, "<button ")
}

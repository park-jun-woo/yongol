//ff:func feature=stml-gen type=test control=sequence
//ff:what Delete 외 액션 버튼에 variant 미포함을 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestUIButtonNoVariantForNonDelete(t *testing.T) {
	page, _ := stmlparser.ParseReader("action-page.html", strings.NewReader(`<main>
  <button data-action="ActivateWorkflow">활성화</button>
</main>`))
	code := GeneratePage(page, "")
	assertNotContains(t, code, `variant="destructive"`)
}

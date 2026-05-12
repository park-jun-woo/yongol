//ff:func feature=stml-gen type=test control=sequence
//ff:what 제약조건 없이 폼 페이지 생성 시 zod 미포함 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePageWithoutConstraints(t *testing.T) {
	page, _ := stmlparser.ParseReader("edit-page.html", strings.NewReader(`<main>
  <div data-action="UpdateItem">
    <input data-field="Name" placeholder="이름" />
    <button type="submit">수정</button>
  </div>
</main>`))

	code := GeneratePage(page, "")
	assertNotContains(t, code, "import { z }")
	assertNotContains(t, code, "zodResolver")
	assertNotContains(t, code, "z.object")
	assertContains(t, code, "const updateItemForm = useForm()")
}

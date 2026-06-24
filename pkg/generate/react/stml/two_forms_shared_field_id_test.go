//ff:func feature=stml-gen type=test control=sequence
//ff:what 동일 페이지 2개 폼이 동명 필드를 가져도 DOM id가 폼 스코프로 유니크함을 검증 (BUG-127)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTwoFormsSharedFieldScopedID(t *testing.T) {
	page, _ := stmlparser.ParseReader("building-detail.html", strings.NewReader(`<main>
  <article data-fetch="GetBuilding">
    <form data-action="UpdateBuilding">
      <input data-field="memo" />
      <button type="submit">수정</button>
    </form>
    <form data-action="CreateBuilding">
      <input data-field="memo" />
      <button type="submit">생성</button>
    </form>
  </article>
</main>`))
	code := GeneratePage(page, "")

	// each form's field id is form-scoped and appears exactly once.
	if n := strings.Count(code, `id="updateBuilding-memo"`); n != 1 {
		t.Errorf(`id="updateBuilding-memo" count = %d, want 1`, n)
	}
	if n := strings.Count(code, `id="createBuilding-memo"`); n != 1 {
		t.Errorf(`id="createBuilding-memo" count = %d, want 1`, n)
	}
	// no bare duplicate id survives.
	if strings.Contains(code, `id="memo"`) {
		t.Errorf(`bare id="memo" must not be emitted: %q`, code)
	}
	// each label's htmlFor matches its own form's input id.
	if !strings.Contains(code, `<label htmlFor="updateBuilding-memo" className="text-sm font-medium">`) {
		t.Errorf("missing update form label htmlFor")
	}
	if !strings.Contains(code, `<label htmlFor="createBuilding-memo" className="text-sm font-medium">`) {
		t.Errorf("missing create form label htmlFor")
	}
	// register/zod keys stay the bare contract field name (unchanged).
	if !strings.Contains(code, `register('memo')`) {
		t.Errorf("register key must stay bare field name 'memo'")
	}
}

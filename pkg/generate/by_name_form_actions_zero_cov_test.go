//ff:func feature=generate type=test control=sequence
//ff:what TestByName_ZeroCov — generate 폼 액션/필드 해석 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestByNameFormActions_ZeroCov(t *testing.T) {
	page := byNameFormPage()
	pages := []stmlparser.PageSpec{page}

	entries := collectFormActionOps(pages)
	if len(entries) == 0 {
		t.Fatalf("collectFormActionOps empty")
	}

	seen := map[string]bool{}
	got := appendPageFormActions(nil, page, seen)
	if len(got) == 0 {
		t.Errorf("appendPageFormActions empty")
	}

	nested := collectNestedFormActions(page.Children, map[string]bool{})
	_ = nested

	fieldless := map[string]bool{}
	collectPageFieldlessActions(page, fieldless)
	if !fieldless["DeleteItem"] {
		t.Errorf("collectPageFieldlessActions missing DeleteItem")
	}

	ae := toActionEntry(page.Actions[0])
	if ae.opID != "CreateItem" || len(ae.fieldNames) != 2 {
		t.Errorf("toActionEntry = %+v", ae)
	}
}

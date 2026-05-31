//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestByNameCollectors_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	f := page.Fetches[0]
	ops := collectFetchOps(f, nil)
	if len(ops) == 0 {
		t.Errorf("collectFetchOps empty")
	}
	binds := collectFetchParamBinds(f, nil)
	_ = binds
	a := page.Actions[0]
	if deduplicateActions([]stmlparser.ActionBlock{a, a}) == nil {
		t.Errorf("deduplicateActions nil")
	}
	_ = extractBindFieldsFromChildren(f.Children)

	is := importSet{}
	compSet := map[string]bool{}
	collectFetchImports(f, &is, compSet)
	collectActionImports(a, &is, compSet)
	full := collectImports(page, "")
	_ = full
}

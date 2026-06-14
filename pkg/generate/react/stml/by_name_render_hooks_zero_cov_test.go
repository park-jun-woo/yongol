//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react/stml 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package stml

import (
	"strings"
	"testing"
)

func TestByNameRenderHooks_ZeroCov(t *testing.T) {
	page := byNameSamplePage(t)
	cons := byNameConstraints()
	noBody := map[string]bool{}
	ppt := map[string]map[string]string{}
	f := page.Fetches[0]
	a := page.Actions[0]

	if s := renderUseQuery(f, ppt); s == "" {
		t.Errorf("renderUseQuery empty")
	}
	if s := renderUseMutation(a, []string{"ListItems"}, nil, false, noBody, ppt, cons, "error"); s == "" {
		t.Errorf("renderUseMutation empty")
	}

	var sb strings.Builder
	renderFetchHooks(f, ppt, &sb)
	is := collectImports(page, "")
	renderPageHooks(page, is, ppt, &sb)
	renderPageMutations(page.Actions, []string{"ListItems"}, buildActionFetchMap(page), cons, false, noBody, ppt, "error", nil, &sb)
	if sb.Len() == 0 {
		t.Errorf("hook renderers produced nothing")
	}

	imports := renderImports(is, DefaultOptions())
	if imports == "" {
		t.Errorf("renderImports empty")
	}

	var jb strings.Builder
	renderPageJSX(page, &jb, noBody, bindCtx{})
	renderPageJSXWithChildren(page.Children, &jb, noBody, bindCtx{})
	if jb.Len() == 0 {
		t.Errorf("renderPageJSX produced nothing")
	}
}

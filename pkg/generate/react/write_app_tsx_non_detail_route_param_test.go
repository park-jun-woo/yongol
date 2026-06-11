//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX non-detail 페이지의 route 파라미터 유도 라우트(필수 세그먼트·액션 전용 optional) 생성 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_NonDetailPageWithRouteParam(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{
			Name:     "templates",
			FileName: "templates.html",
			Fetches: []stml.FetchBlock{
				{OperationID: "ListTemplates"},
				{
					OperationID: "GetTemplate",
					Params:      []stml.ParamBind{{Name: "id", Source: "route.id"}},
				},
			},
		},
		{
			Name:     "webhooks",
			FileName: "webhooks.html",
			Fetches: []stml.FetchBlock{
				{OperationID: "ListWebhooks"},
			},
			Actions: []stml.ActionBlock{
				{
					OperationID: "DeleteWebhook",
					Params:      []stml.ParamBind{{Name: "id", Source: "route.id"}},
				},
			},
		},
	}
	if err := writeAppTSX(dir, pages, nil, "", nil, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// fetch-consumed param → single route with a required segment;
	// the bare base path is no longer emitted (it would render with
	// undefined params — the NaN-call false positive of BUG-112).
	assertContains(t, content, `<Route path="/templates/:id" element={<Templates />} />`)
	assertNotContains(t, content, `<Route path="/templates" element=`)

	// action-only param → optional trailing segment, page stays reachable
	assertContains(t, content, `<Route path="/webhooks/:id?" element={<Webhooks />} />`)
	assertNotContains(t, content, `<Route path="/webhooks" element=`)

	importCount := strings.Count(content, "import Templates from")
	if importCount != 1 {
		t.Errorf("expected 1 Templates import, got %d", importCount)
	}
	importCount = strings.Count(content, "import Webhooks from")
	if importCount != 1 {
		t.Errorf("expected 1 Webhooks import, got %d", importCount)
	}
}

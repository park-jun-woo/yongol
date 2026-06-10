//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX non-detail 페이지의 route 파라미터 이중 라우트 생성 검증

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
	if err := writeAppTSX(dir, pages, nil, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, `<Route path="/templates" element={<Templates />} />`)
	assertContains(t, content, `<Route path="/templates/:id" element={<Templates />} />`)
	assertContains(t, content, `<Route path="/webhooks" element={<Webhooks />} />`)
	assertContains(t, content, `<Route path="/webhooks/:id" element={<Webhooks />} />`)

	importCount := strings.Count(content, "import Templates from")
	if importCount != 1 {
		t.Errorf("expected 1 Templates import, got %d", importCount)
	}
	importCount = strings.Count(content, "import Webhooks from")
	if importCount != 1 {
		t.Errorf("expected 1 Webhooks import, got %d", importCount)
	}
}

//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX data-route 명시적 라우트 우선 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_ExplicitRoute(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{
			Name:     "unit-detail",
			FileName: "unit-detail.html",
			Route:    "/buildings/:buildingId/units/:id",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetUnit",
				Params: []stml.ParamBind{
					{Name: "buildingId", Source: "route.buildingId"},
					{Name: "id", Source: "route.id"},
				},
			}},
		},
	}
	if err := writeAppTSX(dir, pages, nil, "", nil, ""); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, `<Route path="/buildings/:buildingId/units/:id" element={<UnitDetail />} />`)
	if strings.Contains(content, `<Route path="/units/:id"`) {
		t.Error("should not contain filename-inferred route /units/:id when data-route is set")
	}
}

//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX detail 페이지 라우트 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_DetailPage(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "workflows", FileName: "workflows.html"},
		{
			Name:     "workflow-detail",
			FileName: "workflow-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetWorkflow",
				Params: []stml.ParamBind{
					{Name: "WorkflowID", Source: "route.WorkflowID"},
				},
			}},
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

	// workflows is the index target → eager; the detail page is lazy (BUG-133)
	assertContains(t, content, "import Workflows from './pages/workflows'")
	assertContains(t, content, "const WorkflowDetail = lazy(() => import('./pages/workflow-detail'))")
	assertContains(t, content, `<Route path="/workflows" element={<Workflows />} />`)
	// route segment name is derived from the consumed route.* declaration,
	// matching the useParams() destructuring the page emitter produces
	assertContains(t, content, `<Route path="/workflows/:WorkflowID" element={<WorkflowDetail />} />`)
}

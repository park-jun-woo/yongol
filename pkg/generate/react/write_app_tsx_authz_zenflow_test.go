//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX zenflow 전체 — 보호 페이지별 가드 + 공개 로그인/회원가입 + 인덱스/catch-all 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_Authz_ZenflowFull(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "register", FileName: "register.html", Layout: "auth"},
		{Name: "workflows", FileName: "workflows.html"},
		{
			Name:     "workflow-detail",
			FileName: "workflow-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetWorkflow",
				Params:      []stml.ParamBind{{Name: "id", Source: "route.id"}},
			}},
		},
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "templates", FileName: "templates.html"},
		{Name: "webhooks", FileName: "webhooks.html"},
		{Name: "audit-logs", FileName: "audit-logs.html"},
	}
	layouts := []stml.LayoutSpec{
		{
			Name: "app",
			NavItems: []stml.NavItem{
				{Path: "/workflows", Label: "Workflows"},
				{Path: "/templates", Label: "Templates"},
				{Path: "/dashboard", Label: "Dashboard"},
			},
			HasOutlet: true,
		},
		{Name: "auth", HasOutlet: true},
	}
	protected := map[string]bool{
		"workflows.html":       true,
		"workflow-detail.html": true,
		"dashboard.html":       true,
		"templates.html":       true,
		"webhooks.html":        true,
		"audit-logs.html":      true,
	}
	if err := writeAppTSX(dir, pages, layouts, "app", protected); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertNotContains(t, content, "<ProtectedRoute><AppLayout /></ProtectedRoute>")
	assertContains(t, content, `        <Route path="/audit-logs" element={<ProtectedRoute><AuditLogs /></ProtectedRoute>} />`)
	assertContains(t, content, `        <Route path="/dashboard" element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />`)
	assertContains(t, content, `        <Route path="/workflows" element={<ProtectedRoute><Workflows /></ProtectedRoute>} />`)
	assertContains(t, content, `        <Route path="/workflows/:id" element={<ProtectedRoute><WorkflowDetail /></ProtectedRoute>} />`)
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertNotContains(t, content, "<ProtectedRoute><AuthLayout /></ProtectedRoute>")
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
	assertContains(t, content, `        <Route path="/register" element={<Register />} />`)
	assertContains(t, content, "import ProtectedRoute from './components/ProtectedRoute'")
	// first public page in file-name order is login → "/" redirects there
	assertContains(t, content, `<Route path="/" element={<Navigate to="/login" replace />} />`)
	assertContains(t, content, `<Route path="*" element={<Navigate to="/" replace />} />`)

	importCount := strings.Count(content, "import ProtectedRoute from")
	if importCount != 1 {
		t.Errorf("expected 1 ProtectedRoute import, got %d", importCount)
	}
}

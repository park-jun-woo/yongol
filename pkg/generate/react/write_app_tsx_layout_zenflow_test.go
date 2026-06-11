//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX zenflow 전체 레이아웃 그룹핑 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_LayoutGrouping_ZenflowFull(t *testing.T) {
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
	if err := writeAppTSX(dir, pages, layouts, "app", nil, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, `        <Route path="/audit-logs" element={<AuditLogs />} />`)
	assertContains(t, content, `        <Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, `        <Route path="/templates" element={<Templates />} />`)
	assertContains(t, content, `        <Route path="/webhooks" element={<Webhooks />} />`)
	assertContains(t, content, `        <Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `        <Route path="/workflows/:id" element={<WorkflowDetail />} />`)
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
	assertContains(t, content, `        <Route path="/register" element={<Register />} />`)
}

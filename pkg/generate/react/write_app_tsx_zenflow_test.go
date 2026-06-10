//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX zenflow 전체 페이지 라우트 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_FullZenflow(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html"},
		{Name: "register", FileName: "register.html"},
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
	if err := writeAppTSX(dir, pages, nil, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, `<Route path="/audit-logs" element={<AuditLogs />} />`)
	assertContains(t, content, `<Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, `<Route path="/login" element={<Login />} />`)
	assertContains(t, content, `<Route path="/register" element={<Register />} />`)
	assertContains(t, content, `<Route path="/templates" element={<Templates />} />`)
	assertContains(t, content, `<Route path="/webhooks" element={<Webhooks />} />`)
	assertContains(t, content, `<Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `<Route path="/workflows/:id" element={<WorkflowDetail />} />`)
	assertContains(t, content, "import { Routes, Route, Navigate } from 'react-router-dom'")
}

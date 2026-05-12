//ff:func feature=gen-react type=test
//ff:what writeAppTSX 라우트 생성 로직 테스트

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_NoPages_Placeholder(t *testing.T) {
	dir := t.TempDir()
	if err := writeAppTSX(dir, nil, nil, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)
	if !strings.Contains(content, "yongol scaffolded frontend") {
		t.Error("expected placeholder content for empty pages")
	}
}

func TestWriteAppTSX_BasicPages(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html"},
		{Name: "register", FileName: "register.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}
	if err := writeAppTSX(dir, pages, nil, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// Check imports
	assertContains(t, content, "import Login from './pages/login'")
	assertContains(t, content, "import Register from './pages/register'")
	assertContains(t, content, "import Dashboard from './pages/dashboard'")

	// Check routes
	assertContains(t, content, `<Route path="/login" element={<Login />} />`)
	assertContains(t, content, `<Route path="/register" element={<Register />} />`)
	assertContains(t, content, `<Route path="/dashboard" element={<Dashboard />} />`)
}

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
	if err := writeAppTSX(dir, pages, nil, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import Workflows from './pages/workflows'")
	assertContains(t, content, "import WorkflowDetail from './pages/workflow-detail'")
	assertContains(t, content, `<Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `<Route path="/workflows/:id" element={<WorkflowDetail />} />`)
}

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
	if err := writeAppTSX(dir, pages, nil, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// Routes are sorted alphabetically by path
	assertContains(t, content, `<Route path="/audit-logs" element={<AuditLogs />} />`)
	assertContains(t, content, `<Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, `<Route path="/login" element={<Login />} />`)
	assertContains(t, content, `<Route path="/register" element={<Register />} />`)
	assertContains(t, content, `<Route path="/templates" element={<Templates />} />`)
	assertContains(t, content, `<Route path="/webhooks" element={<Webhooks />} />`)
	assertContains(t, content, `<Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `<Route path="/workflows/:id" element={<WorkflowDetail />} />`)

	// Verify react-router-dom import
	assertContains(t, content, "import { Routes, Route } from 'react-router-dom'")
}

func TestKebabToPascal(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"login", "Login"},
		{"workflow-detail", "WorkflowDetail"},
		{"audit-logs", "AuditLogs"},
		{"my-long-page-name", "MyLongPageName"},
		{"dashboard", "Dashboard"},
	}
	for _, tt := range tests {
		got := kebabToPascal(tt.in)
		if got != tt.want {
			t.Errorf("kebabToPascal(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestNaivePluralize(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"workflow", "workflows"},
		{"bus", "buses"},
		{"match", "matches"},
		{"box", "boxes"},
		{"buzz", "buzzes"},
		{"dish", "dishes"},
		{"template", "templates"},
	}
	for _, tt := range tests {
		got := naivePluralize(tt.in)
		if got != tt.want {
			t.Errorf("naivePluralize(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestPageHasRouteParam(t *testing.T) {
	noRoute := stml.PageSpec{
		Name:     "login",
		FileName: "login.html",
	}
	if pageHasRouteParam(noRoute) {
		t.Error("expected no route param for login page")
	}

	withFetchRoute := stml.PageSpec{
		Name:     "workflow-detail",
		FileName: "workflow-detail.html",
		Fetches: []stml.FetchBlock{{
			Params: []stml.ParamBind{{Name: "id", Source: "route.id"}},
		}},
	}
	if !pageHasRouteParam(withFetchRoute) {
		t.Error("expected route param for page with route.id in fetch")
	}

	withActionRoute := stml.PageSpec{
		Name:     "room-edit",
		FileName: "room-edit.html",
		Actions: []stml.ActionBlock{{
			Params: []stml.ParamBind{{Name: "RoomID", Source: "route.RoomID"}},
		}},
	}
	if !pageHasRouteParam(withActionRoute) {
		t.Error("expected route param for page with route.RoomID in action")
	}
}

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
	if err := writeAppTSX(dir, pages, nil, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// templates: non-detail page with route param → two routes
	assertContains(t, content, `<Route path="/templates" element={<Templates />} />`)
	assertContains(t, content, `<Route path="/templates/:id" element={<Templates />} />`)

	// webhooks: non-detail page with route param → two routes
	assertContains(t, content, `<Route path="/webhooks" element={<Webhooks />} />`)
	assertContains(t, content, `<Route path="/webhooks/:id" element={<Webhooks />} />`)

	// Import should appear only once per component
	importCount := strings.Count(content, "import Templates from")
	if importCount != 1 {
		t.Errorf("expected 1 Templates import, got %d", importCount)
	}
	importCount = strings.Count(content, "import Webhooks from")
	if importCount != 1 {
		t.Errorf("expected 1 Webhooks import, got %d", importCount)
	}
}

func TestBuildRoutes_Sorted(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "zebra", FileName: "zebra.html"},
		{Name: "alpha", FileName: "alpha.html"},
		{Name: "mid", FileName: "mid.html"},
	}
	routes := buildRoutes(pages, "")
	if len(routes) != 3 {
		t.Fatalf("got %d routes, want 3", len(routes))
	}
	if routes[0].Path != "/alpha" {
		t.Errorf("routes[0].Path = %q, want /alpha", routes[0].Path)
	}
	if routes[1].Path != "/mid" {
		t.Errorf("routes[1].Path = %q, want /mid", routes[1].Path)
	}
	if routes[2].Path != "/zebra" {
		t.Errorf("routes[2].Path = %q, want /zebra", routes[2].Path)
	}
}

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
	if err := writeAppTSX(dir, pages, nil, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// Explicit route must override filename-based inference
	assertContains(t, content, `<Route path="/buildings/:buildingId/units/:id" element={<UnitDetail />} />`)
	// Must NOT have the inferred route
	if strings.Contains(content, `<Route path="/units/:id"`) {
		t.Error("should not contain filename-inferred route /units/:id when data-route is set")
	}
}

func TestPageToRoutes_ExplicitRouteOverridesDetail(t *testing.T) {
	p := stml.PageSpec{
		Name:     "unit-detail",
		FileName: "unit-detail.html",
		Route:    "/buildings/:buildingId/units/:id",
	}
	routes := pageToRoutes(p)
	if len(routes) != 1 {
		t.Fatalf("got %d routes, want 1", len(routes))
	}
	if routes[0].Path != "/buildings/:buildingId/units/:id" {
		t.Errorf("Path = %q, want /buildings/:buildingId/units/:id", routes[0].Path)
	}
	if routes[0].ComponentName != "UnitDetail" {
		t.Errorf("ComponentName = %q, want UnitDetail", routes[0].ComponentName)
	}
}

func assertContains(t *testing.T, content, want string) {
	t.Helper()
	if !strings.Contains(content, want) {
		t.Errorf("output missing expected substring:\n  want: %q\n  got:\n%s", want, content)
	}
}

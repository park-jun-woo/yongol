//ff:func feature=gen-react type=test
//ff:what writeLayoutTSX 레이아웃 TSX 생성 + App.tsx 레이아웃 route 그룹핑 테스트

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteLayoutTSX_AppLayout(t *testing.T) {
	dir := t.TempDir()
	layout := stml.LayoutSpec{
		Name: "app",
		File: "layouts/app.html",
		NavItems: []stml.NavItem{
			{Path: "/workflows", Label: "Workflows"},
			{Path: "/dashboard", Label: "Dashboard"},
		},
		HasOutlet: true,
	}
	if err := writeLayoutTSX(dir, layout); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "layouts", "AppLayout.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { Link, Outlet } from 'react-router-dom'")
	assertContains(t, content, "export default function AppLayout()")
	assertContains(t, content, `<Link to="/workflows">Workflows</Link>`)
	assertContains(t, content, `<Link to="/dashboard">Dashboard</Link>`)
	assertContains(t, content, "<Outlet />")
	assertContains(t, content, "<nav>")
}

func TestWriteLayoutTSX_AuthLayout_NoNav(t *testing.T) {
	dir := t.TempDir()
	layout := stml.LayoutSpec{
		Name:      "auth",
		File:      "layouts/auth.html",
		HasOutlet: true,
	}
	if err := writeLayoutTSX(dir, layout); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "layouts", "AuthLayout.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { Outlet } from 'react-router-dom'")
	assertContains(t, content, "export default function AuthLayout()")
	assertContains(t, content, "<Outlet />")
	assertNotContains(t, content, "<nav>")
	assertNotContains(t, content, "Link")
}

func TestWriteLayoutsTSX_Multiple(t *testing.T) {
	dir := t.TempDir()
	layouts := []stml.LayoutSpec{
		{
			Name: "app",
			NavItems: []stml.NavItem{
				{Path: "/workflows", Label: "Workflows"},
			},
			HasOutlet: true,
		},
		{
			Name:      "auth",
			HasOutlet: true,
		},
	}
	if err := writeLayoutsTSX(dir, layouts); err != nil {
		t.Fatal(err)
	}

	// Both files exist
	for _, name := range []string{"AppLayout.tsx", "AuthLayout.tsx"} {
		if _, err := os.Stat(filepath.Join(dir, "layouts", name)); err != nil {
			t.Errorf("expected %s to exist: %v", name, err)
		}
	}
}

func TestLayoutComponentName(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"app", "AppLayout"},
		{"auth", "AuthLayout"},
		{"main-nav", "MainNavLayout"},
		{"admin-panel", "AdminPanelLayout"},
	}
	for _, tt := range tests {
		got := layoutComponentName(tt.in)
		if got != tt.want {
			t.Errorf("layoutComponentName(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestWriteAppTSX_LayoutGrouping(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "register", FileName: "register.html", Layout: "auth"},
		{Name: "workflows", FileName: "workflows.html", Layout: "app"},
		{Name: "dashboard", FileName: "dashboard.html", Layout: "app"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", HasOutlet: true, NavItems: []stml.NavItem{{Path: "/workflows", Label: "Workflows"}}},
		{Name: "auth", HasOutlet: true},
	}
	if err := writeAppTSX(dir, pages, layouts, ""); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// Layout imports
	assertContains(t, content, "import AppLayout from './layouts/AppLayout'")
	assertContains(t, content, "import AuthLayout from './layouts/AuthLayout'")

	// Layout wrapper routes
	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, "<Route element={<AuthLayout />}>")

	// Child routes inside wrappers (indented deeper)
	assertContains(t, content, `        <Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `        <Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
	assertContains(t, content, `        <Route path="/register" element={<Register />} />`)

	// No flat routes (all pages have layouts)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "<Route path=") && !strings.Contains(line, "        ") {
			// Allow lines that are inside a layout group (8 spaces indent)
			if strings.HasPrefix(line, "      <Route path=") && !strings.HasPrefix(line, "        ") {
				t.Errorf("unexpected flat route found: %s", line)
			}
		}
	}
}

func TestWriteAppTSX_DefaultLayout(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "workflows", FileName: "workflows.html"},       // no explicit layout
		{Name: "dashboard", FileName: "dashboard.html"},        // no explicit layout
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", HasOutlet: true},
		{Name: "auth", HasOutlet: true},
	}
	// defaultLayout = "app" → workflows and dashboard go under AppLayout
	if err := writeAppTSX(dir, pages, layouts, "app"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// workflows and dashboard should be grouped under AppLayout
	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, `        <Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `        <Route path="/dashboard" element={<Dashboard />} />`)

	// login should be grouped under AuthLayout
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
}

func TestWriteAppTSX_NoLayoutNoDefault_FlatRoutes(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}
	// No layouts, no defaultLayout → flat routes (current behavior)
	if err := writeAppTSX(dir, pages, nil, ""); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, `      <Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, `      <Route path="/login" element={<Login />} />`)
	assertNotContains(t, content, "Layout")
}

func TestWriteAppTSX_MixedLayoutAndFlat(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "about", FileName: "about.html"}, // no layout, no default
		{Name: "workflows", FileName: "workflows.html", Layout: "app"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", HasOutlet: true},
		{Name: "auth", HasOutlet: true},
	}
	// No defaultLayout → "about" stays flat
	if err := writeAppTSX(dir, pages, layouts, ""); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// Layout groups
	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, "<Route element={<AuthLayout />}>")

	// Flat route for about (6 spaces, not 8)
	assertContains(t, content, `      <Route path="/about" element={<About />} />`)

	// Layout imports present
	assertContains(t, content, "import AppLayout from './layouts/AppLayout'")
	assertContains(t, content, "import AuthLayout from './layouts/AuthLayout'")
}

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
	// defaultLayout = "app" → pages without explicit layout go under AppLayout
	if err := writeAppTSX(dir, pages, layouts, "app"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// AppLayout group should contain: audit-logs, dashboard, templates, webhooks, workflows, workflows/:id
	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, `        <Route path="/audit-logs" element={<AuditLogs />} />`)
	assertContains(t, content, `        <Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, `        <Route path="/templates" element={<Templates />} />`)
	assertContains(t, content, `        <Route path="/webhooks" element={<Webhooks />} />`)
	assertContains(t, content, `        <Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `        <Route path="/workflows/:id" element={<WorkflowDetail />} />`)

	// AuthLayout group should contain: login, register
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
	assertContains(t, content, `        <Route path="/register" element={<Register />} />`)
}

func TestLayoutImports_LinkAndOutlet(t *testing.T) {
	layout := stml.LayoutSpec{
		NavItems:  []stml.NavItem{{Path: "/foo", Label: "Foo"}},
		HasOutlet: true,
	}
	imports := layoutImports(layout)
	if len(imports) != 2 || imports[0] != "Link" || imports[1] != "Outlet" {
		t.Errorf("expected [Link, Outlet], got %v", imports)
	}
}

func TestLayoutImports_OutletOnly(t *testing.T) {
	layout := stml.LayoutSpec{HasOutlet: true}
	imports := layoutImports(layout)
	if len(imports) != 1 || imports[0] != "Outlet" {
		t.Errorf("expected [Outlet], got %v", imports)
	}
}

func TestLayoutImports_Empty(t *testing.T) {
	layout := stml.LayoutSpec{}
	imports := layoutImports(layout)
	if len(imports) != 0 {
		t.Errorf("expected empty imports, got %v", imports)
	}
}

func assertNotContains(t *testing.T, content, unwanted string) {
	t.Helper()
	if strings.Contains(content, unwanted) {
		t.Errorf("output should not contain %q but does:\n%s", unwanted, content)
	}
}

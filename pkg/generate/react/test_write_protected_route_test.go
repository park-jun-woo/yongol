//ff:func feature=gen-react type=test
//ff:what ProtectedRoute 컴포넌트 생성 + App.tsx 인증 가드 래핑 + api.ts JWT middleware 테스트

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteProtectedRoute(t *testing.T) {
	dir := t.TempDir()
	if err := writeProtectedRoute(dir); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "components", "ProtectedRoute.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { Navigate } from 'react-router-dom'")
	assertContains(t, content, "export default function ProtectedRoute")
	assertContains(t, content, "localStorage.getItem('access_token')")
	assertContains(t, content, `<Navigate to="/login" replace />`)
}

func TestWriteAppTSX_Authz_LayoutGrouping(t *testing.T) {
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
	// hasAuthz = true → AppLayout wrapped, AuthLayout not wrapped
	if err := writeAppTSX(dir, pages, layouts, "", true); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// ProtectedRoute import
	assertContains(t, content, "import ProtectedRoute from './components/ProtectedRoute'")

	// AppLayout wrapped with ProtectedRoute
	assertContains(t, content, "<Route element={<ProtectedRoute><AppLayout /></ProtectedRoute>}>")

	// AuthLayout NOT wrapped
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertNotContains(t, content, "<ProtectedRoute><AuthLayout /></ProtectedRoute>")

	// Child routes still present
	assertContains(t, content, `        <Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
}

func TestWriteAppTSX_Authz_FlatRoutes(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "about", FileName: "about.html"},
		{Name: "settings", FileName: "settings.html"},
	}
	// hasAuthz = true, no layouts → flat routes wrapped with ProtectedRoute
	if err := writeAppTSX(dir, pages, nil, "", true); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import ProtectedRoute from './components/ProtectedRoute'")
	assertContains(t, content, `<Route path="/about" element={<ProtectedRoute><About /></ProtectedRoute>} />`)
	assertContains(t, content, `<Route path="/settings" element={<ProtectedRoute><Settings /></ProtectedRoute>} />`)
}

func TestWriteAppTSX_Authz_MixedLayoutAndFlat(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "about", FileName: "about.html"}, // no layout
		{Name: "workflows", FileName: "workflows.html", Layout: "app"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", HasOutlet: true},
		{Name: "auth", HasOutlet: true},
	}
	if err := writeAppTSX(dir, pages, layouts, "", true); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// AppLayout wrapped
	assertContains(t, content, "<Route element={<ProtectedRoute><AppLayout /></ProtectedRoute>}>")
	// AuthLayout NOT wrapped
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	// Flat route wrapped
	assertContains(t, content, `<Route path="/about" element={<ProtectedRoute><About /></ProtectedRoute>} />`)
}

func TestWriteAppTSX_NoAuthz_NoProtectedRoute(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "workflows", FileName: "workflows.html", Layout: "app"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", HasOutlet: true},
		{Name: "auth", HasOutlet: true},
	}
	// hasAuthz = false → no ProtectedRoute
	if err := writeAppTSX(dir, pages, layouts, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertNotContains(t, content, "ProtectedRoute")
	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, "<Route element={<AuthLayout />}>")
}

func TestWriteAPIClient_Authz_JWTMiddleware(t *testing.T) {
	dir := t.TempDir()
	// hasAuthz = true → JWT middleware emitted
	if err := writeAPIClient(dir, nil, true); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "client.use({")
	assertContains(t, content, "async onRequest({ request })")
	assertContains(t, content, "localStorage.getItem('access_token')")
	assertContains(t, content, "request.headers.set('Authorization', `Bearer ${token}`)")
}

func TestWriteAPIClient_NoAuthz_NoMiddleware(t *testing.T) {
	dir := t.TempDir()
	// hasAuthz = false → no middleware
	if err := writeAPIClient(dir, nil, false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "lib", "api.ts"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertNotContains(t, content, "client.use")
	assertNotContains(t, content, "onRequest")
	assertNotContains(t, content, "Authorization")
}

func TestIsAuthLayout(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"auth", true},
		{"app", false},
		{"admin", false},
		{"", false},
	}
	for _, tt := range tests {
		got := isAuthLayout(tt.name)
		if got != tt.want {
			t.Errorf("isAuthLayout(%q) = %v, want %v", tt.name, got, tt.want)
		}
	}
}

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
	// hasAuthz = true, defaultLayout = "app"
	if err := writeAppTSX(dir, pages, layouts, "app", true); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// AppLayout wrapped with ProtectedRoute
	assertContains(t, content, "<Route element={<ProtectedRoute><AppLayout /></ProtectedRoute>}>")
	// All app pages under protected wrapper
	assertContains(t, content, `        <Route path="/audit-logs" element={<AuditLogs />} />`)
	assertContains(t, content, `        <Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, `        <Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `        <Route path="/workflows/:id" element={<WorkflowDetail />} />`)

	// AuthLayout NOT wrapped
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertNotContains(t, content, "<ProtectedRoute><AuthLayout /></ProtectedRoute>")
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
	assertContains(t, content, `        <Route path="/register" element={<Register />} />`)

	// ProtectedRoute import present
	assertContains(t, content, "import ProtectedRoute from './components/ProtectedRoute'")

	// ProtectedRoute import appears exactly once
	importCount := strings.Count(content, "import ProtectedRoute from")
	if importCount != 1 {
		t.Errorf("expected 1 ProtectedRoute import, got %d", importCount)
	}
}

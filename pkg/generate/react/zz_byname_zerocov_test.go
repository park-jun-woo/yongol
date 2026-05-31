//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// TestByNamePureHelpers_ZeroCov exercises small pure helpers by name.
func TestByNamePureHelpers_ZeroCov(t *testing.T) {
	if got := buildCNArgs("base", []string{"v"}, []string{"s"}); len(got) == 0 {
		t.Errorf("buildCNArgs empty")
	}
	_ = buildCNArgs("", nil, nil)
	if got := buildDestructParams([]string{"v"}, []string{"s"}, "default", "md"); len(got) == 0 {
		t.Errorf("buildDestructParams empty")
	}
	_ = buildDestructParams(nil, nil, "", "")

	if got := extractPathParams("/api/items/{id}/sub/{subId}"); len(got) != 2 {
		t.Errorf("extractPathParams = %v", got)
	}
	_ = extractPathParams("/api/items")

	for _, tag := range []string{"button", "input", "select", "textarea", "form", "table", "label", "a", "div"} {
		_ = htmlAttrsType(tag)
		_ = inferHTMLElement(tag)
	}
	for _, name := range []string{"Button", "Input", "Select", "Textarea", "Form", "Table", "Label", "Link", "Card"} {
		_ = inferHTMLTag(name)
	}
	if lcFirst("ItemID") == "" {
		t.Errorf("lcFirst empty")
	}

	dspec := &design.DesignSpec{Colors: map[string]string{"primary": "#fff"}}
	if got := designColor(dspec, "primary", "def"); got != "#fff" {
		t.Errorf("designColor = %q", got)
	}
	_ = designColor(dspec, "missing", "def")
	_ = designColor(nil, "primary", "def")

	set := buildLayoutSet([]stml.LayoutSpec{{Name: "app"}, {Name: "auth"}})
	if !set["app"] {
		t.Errorf("buildLayoutSet missing app")
	}

	if !hasRouteSource([]stml.ParamBind{{Name: "ID", Source: "route.ID"}}) {
		t.Errorf("hasRouteSource = false")
	}
	_ = hasRouteSource([]stml.ParamBind{{Name: "X", Source: "state.X"}})

	grp := groupRoutesByLayout([]stmlRoute{{Path: "/a", Layout: "app"}, {Path: "/b", Layout: "app"}})
	if len(grp["app"]) != 2 {
		t.Errorf("groupRoutesByLayout = %v", grp)
	}
}

// TestByNameEndpoints_ZeroCov exercises collectEndpoints / appendOperations by name.
func TestByNameEndpoints_ZeroCov(t *testing.T) {
	doc := &openapi3.T{
		Paths: openapi3.NewPaths(
			openapi3.WithPath("/api/items/{id}", &openapi3.PathItem{
				Get:  &openapi3.Operation{OperationID: "GetItem"},
				Post: &openapi3.Operation{OperationID: "CreateItem"},
			}),
		),
	}
	eps := collectEndpoints(doc)
	if len(eps) != 2 {
		t.Fatalf("collectEndpoints = %d, want 2", len(eps))
	}

	pi := &openapi3.PathItem{Get: &openapi3.Operation{OperationID: "X"}}
	got := appendOperations(nil, "/x", pi)
	if len(got) != 1 {
		t.Errorf("appendOperations = %d", len(got))
	}

	var b strings.Builder
	writeApiClientEntry(&b, eps[0])
	if b.Len() == 0 {
		t.Errorf("writeApiClientEntry empty")
	}
}

// TestByNameRenderComponents_ZeroCov exercises render*Component by name.
func TestByNameRenderComponents_ZeroCov(t *testing.T) {
	simple := renderSimpleComponent("Input", design.ComponentToken{Base: "px-2"})
	if !strings.Contains(simple, "Input") {
		t.Errorf("renderSimpleComponent missing name")
	}
	_ = renderSimpleComponent("Card", design.ComponentToken{})

	variant := renderVariantComponent("Button", design.ComponentToken{
		Base:           "px-4",
		Variants:       map[string]string{"primary": "bg-blue", "ghost": "bg-none"},
		Sizes:          map[string]string{"sm": "text-sm", "lg": "text-lg"},
		DefaultVariant: "primary",
		DefaultSize:    "sm",
	})
	if !strings.Contains(variant, "Button") {
		t.Errorf("renderVariantComponent missing name")
	}
}

// TestByNameVariantWriters_ZeroCov exercises writeVariant* by name.
func TestByNameVariantWriters_ZeroCov(t *testing.T) {
	vk := []string{"primary", "ghost"}
	sk := []string{"sm", "lg"}
	tok := design.ComponentToken{
		Variants: map[string]string{"primary": "bg-blue", "ghost": "bg-none"},
		Sizes:    map[string]string{"sm": "text-sm", "lg": "text-lg"},
	}
	var b strings.Builder
	writeVariantTypes(&b, vk, sk)
	writeVariantProps(&b, "Button", "button", vk, sk)
	writeVariantRecords(&b, vk, sk, tok)
	if b.Len() == 0 {
		t.Errorf("variant writers produced nothing")
	}
}

// TestByNameColorAndTailwind_ZeroCov exercises color/tailwind writers by name.
func TestByNameColorAndTailwind_ZeroCov(t *testing.T) {
	var b strings.Builder
	writeColorToken(&b, "primary", "#fff", "#000")
	writeExtraDesignColors(&b, map[string]string{"brand": "#123", "accent2": "#456"})
	if b.Len() == 0 {
		t.Errorf("color writers empty")
	}

	theme := &manifest.FrontendTheme{Radius: "0.5rem"}
	dspec := &design.DesignSpec{
		Rounded: map[string]string{"sm": "2px", "lg": "8px"},
		Spacing: map[string]string{"xs": "4px", "xl": "32px"},
	}
	var b2 strings.Builder
	writeTailwindBorderRadius(&b2, theme, dspec)
	writeTailwindSpacing(&b2, dspec)
	if b2.Len() == 0 {
		t.Errorf("tailwind writers empty")
	}
}

// TestByNameRouteWriters_ZeroCov exercises App.tsx route writers by name.
func TestByNameRouteWriters_ZeroCov(t *testing.T) {
	routes := []stmlRoute{
		{Path: "/", ComponentName: "Home", ImportPath: "./pages/home", Layout: ""},
		{Path: "/items", ComponentName: "Items", ImportPath: "./pages/items", Layout: "app"},
	}
	layoutSet := map[string]bool{"app": true}

	app := renderAppTSX(routes, layoutSet, true)
	if !strings.Contains(app, "Routes") {
		t.Errorf("renderAppTSX missing Routes")
	}
	_ = renderAppTSX(routes, layoutSet, false)

	var sb strings.Builder
	writeLayoutImports(&sb, []string{"app", "auth"})
	writePageImports(&sb, routes)
	writeFlatRoutes(&sb, []stmlRoute{routes[0]}, true)
	writeLayoutGroupRoutes(&sb, "app", []stmlRoute{routes[1]}, true)
	writeAuthzMiddleware(&sb)
	if sb.Len() == 0 {
		t.Errorf("route writers empty")
	}
}

// TestByNameLayoutTSX_ZeroCov exercises renderLayoutTSX by name.
func TestByNameLayoutTSX_ZeroCov(t *testing.T) {
	layout := stml.LayoutSpec{
		Name:      "app",
		NavItems:  []stml.NavItem{{Path: "/home", Label: "Home"}},
		HasOutlet: true,
	}
	out := renderLayoutTSX("AppLayout", layout)
	if !strings.Contains(out, "AppLayout") {
		t.Errorf("renderLayoutTSX missing name")
	}
}

// TestByNameFileWriters_ZeroCov exercises file-emitting helpers by name.
func TestByNameFileWriters_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := writeLibUtils(dir); err != nil {
		t.Fatalf("writeLibUtils: %v", err)
	}
	if err := writeAppTSXPlaceholder(dir); err != nil {
		t.Fatalf("writeAppTSXPlaceholder: %v", err)
	}

	uiDir := filepath.Join(dir, "ui")
	if err := os.MkdirAll(uiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	comps := map[string]design.ComponentToken{
		"Input":  {Base: "px-2"},
		"Button": {Variants: map[string]string{"primary": "bg-blue"}},
	}
	if err := writeDesignComponents(uiDir, comps); err != nil {
		t.Fatalf("writeDesignComponents: %v", err)
	}
}

//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestParseLayoutAppNavAndOutlet — app 레이아웃 data-nav + data-outlet 추출 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseLayoutAppNavAndOutlet(t *testing.T) {
	input := `<div>
  <nav>
    <a data-nav="/workflows">Workflows</a>
    <a data-nav="/templates">Templates</a>
    <a data-nav="/dashboard">Dashboard</a>
  </nav>
  <slot data-outlet />
</div>`

	layout, diags := ParseLayoutReader("app.html", "layouts/app.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	if layout.Name != "app" {
		t.Errorf("Name = %q, want %q", layout.Name, "app")
	}
	if layout.File != "layouts/app.html" {
		t.Errorf("File = %q, want %q", layout.File, "layouts/app.html")
	}
	if !layout.HasOutlet {
		t.Error("HasOutlet = false, want true")
	}
	if len(layout.NavItems) != 3 {
		t.Fatalf("NavItems = %d, want 3", len(layout.NavItems))
	}

	wantNav := []NavItem{
		{Path: "/workflows", Label: "Workflows"},
		{Path: "/templates", Label: "Templates"},
		{Path: "/dashboard", Label: "Dashboard"},
	}
	for i, want := range wantNav {
		got := layout.NavItems[i]
		if got.Path != want.Path {
			t.Errorf("NavItems[%d].Path = %q, want %q", i, got.Path, want.Path)
		}
		if got.Label != want.Label {
			t.Errorf("NavItems[%d].Label = %q, want %q", i, got.Label, want.Label)
		}
	}
}

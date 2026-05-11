//ff:func feature=stml-parse type=test control=sequence
//ff:what 레이아웃 파싱 테스트 — data-nav, slot data-outlet 추출 검증

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

func TestParseLayoutAuthNoNav(t *testing.T) {
	input := `<div class="auth-layout">
  <slot data-outlet />
</div>`

	layout, diags := ParseLayoutReader("auth.html", "layouts/auth.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	if layout.Name != "auth" {
		t.Errorf("Name = %q, want %q", layout.Name, "auth")
	}
	if !layout.HasOutlet {
		t.Error("HasOutlet = false, want true")
	}
	if len(layout.NavItems) != 0 {
		t.Errorf("NavItems = %d, want 0", len(layout.NavItems))
	}
}

func TestParseLayoutNoOutlet(t *testing.T) {
	input := `<div>
  <nav>
    <a data-nav="/home">Home</a>
  </nav>
</div>`

	layout, diags := ParseLayoutReader("bare.html", "layouts/bare.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	if layout.HasOutlet {
		t.Error("HasOutlet = true, want false")
	}
	if len(layout.NavItems) != 1 {
		t.Fatalf("NavItems = %d, want 1", len(layout.NavItems))
	}
	if layout.NavItems[0].Path != "/home" {
		t.Errorf("NavItems[0].Path = %q, want %q", layout.NavItems[0].Path, "/home")
	}
}

//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseLayoutAuthNoNav — nav 없는 auth 레이아웃 파싱 검증

package stml

import (
	"strings"
	"testing"
)

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

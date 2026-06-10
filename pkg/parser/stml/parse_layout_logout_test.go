//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseLayoutLogout — data-logout 수집 검증 (op 값/값 없음/부재/첫 선언 우선)

package stml

import (
	"strings"
	"testing"
)

func TestParseLayoutLogout(t *testing.T) {
	parse := func(t *testing.T, input string) LayoutSpec {
		t.Helper()
		layout, diags := ParseLayoutReader("app.html", "layouts/app.html", strings.NewReader(input))
		if len(diags) > 0 {
			t.Fatal(diags)
		}
		return layout
	}

	t.Run("with operationId", func(t *testing.T) {
		layout := parse(t, `<div>
  <nav>
    <a data-nav="dashboard">대시보드</a>
    <button data-logout="Logout">로그아웃</button>
  </nav>
  <slot data-outlet />
</div>`)
		if layout.Logout == nil {
			t.Fatal("Logout = nil, want non-nil")
		}
		if layout.Logout.OperationID != "Logout" {
			t.Errorf("Logout.OperationID = %q, want %q", layout.Logout.OperationID, "Logout")
		}
		if layout.Logout.Label != "로그아웃" {
			t.Errorf("Logout.Label = %q, want %q", layout.Logout.Label, "로그아웃")
		}
		if len(layout.NavItems) != 1 || layout.NavItems[0].Path != "dashboard" {
			t.Errorf("NavItems = %+v, want single page-name entry", layout.NavItems)
		}
	})

	t.Run("valueless", func(t *testing.T) {
		layout := parse(t, `<div>
  <nav><button data-logout>Sign out</button></nav>
  <slot data-outlet />
</div>`)
		if layout.Logout == nil {
			t.Fatal("Logout = nil, want non-nil")
		}
		if layout.Logout.OperationID != "" {
			t.Errorf("Logout.OperationID = %q, want empty", layout.Logout.OperationID)
		}
		if layout.Logout.Label != "Sign out" {
			t.Errorf("Logout.Label = %q, want %q", layout.Logout.Label, "Sign out")
		}
	})

	t.Run("absent", func(t *testing.T) {
		layout := parse(t, `<div><slot data-outlet /></div>`)
		if layout.Logout != nil {
			t.Errorf("Logout = %+v, want nil", layout.Logout)
		}
	})

	t.Run("first occurrence wins", func(t *testing.T) {
		layout := parse(t, `<div>
  <button data-logout="Logout">first</button>
  <button data-logout="SignOut">second</button>
</div>`)
		if layout.Logout == nil {
			t.Fatal("Logout = nil, want non-nil")
		}
		if layout.Logout.OperationID != "Logout" {
			t.Errorf("Logout.OperationID = %q, want first occurrence %q", layout.Logout.OperationID, "Logout")
		}
		if layout.Logout.Label != "first" {
			t.Errorf("Logout.Label = %q, want %q", layout.Logout.Label, "first")
		}
	})
}

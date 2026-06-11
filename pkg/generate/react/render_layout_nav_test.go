//ff:func feature=gen-react type=test control=sequence
//ff:what renderLayoutNav — nav 링크/로그아웃 버튼/빈 블록 생략/라벨 폴백 방출 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderLayoutNav(t *testing.T) {
	patterns := map[string]string{"dashboard": "/dashboard"}

	t.Run("links and logout button", func(t *testing.T) {
		var sb strings.Builder
		layout := stml.LayoutSpec{
			NavItems: []stml.NavItem{{Path: "dashboard", Label: "대시보드"}},
			Logout:   &stml.LogoutSpec{OperationID: "Logout", Label: "로그아웃"},
		}
		renderLayoutNav(&sb, layout, patterns, true, nil)
		out := sb.String()
		assertContains(t, out, "<nav>")
		assertContains(t, out, `<Link to="/dashboard">대시보드</Link>`)
		assertContains(t, out, "<button onClick={handleLogout}>로그아웃</button>")
		assertContains(t, out, "</nav>")
	})

	t.Run("empty label falls back to Logout", func(t *testing.T) {
		var sb strings.Builder
		layout := stml.LayoutSpec{Logout: &stml.LogoutSpec{}}
		renderLayoutNav(&sb, layout, nil, true, nil)
		assertContains(t, sb.String(), "<button onClick={handleLogout}>Logout</button>")
	})

	t.Run("logout suppressed without emission gate", func(t *testing.T) {
		var sb strings.Builder
		layout := stml.LayoutSpec{
			NavItems: []stml.NavItem{{Path: "/home", Label: "Home"}},
			Logout:   &stml.LogoutSpec{OperationID: "Logout"},
		}
		renderLayoutNav(&sb, layout, nil, false, nil)
		out := sb.String()
		assertContains(t, out, `<Link to="/home">Home</Link>`)
		assertNotContains(t, out, "handleLogout")
	})

	t.Run("nothing without nav items or logout", func(t *testing.T) {
		var sb strings.Builder
		renderLayoutNav(&sb, stml.LayoutSpec{}, nil, false, nil)
		if sb.Len() != 0 {
			t.Errorf("expected empty output, got %q", sb.String())
		}
	})
}

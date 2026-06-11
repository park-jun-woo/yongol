//ff:func feature=gen-react type=test control=sequence
//ff:what renderLayoutOutlet — 메뉴/동적 crumb 4분면별 Breadcrumb·Outlet 방출 검증

package react

import (
	"strings"
	"testing"
)

func TestRenderLayoutOutlet(t *testing.T) {
	render := func(hasMenu, dynamicCrumb bool) string {
		var sb strings.Builder
		renderLayoutOutlet(&sb, hasMenu, dynamicCrumb)
		return sb.String()
	}

	t.Run("menu + dynamic crumb wires label prop and Outlet context", func(t *testing.T) {
		got := render(true, true)
		if got != "      <Breadcrumb label={crumbLabel} />\n      <Outlet context={{ setCrumbLabel }} />\n" {
			t.Errorf("got:\n%s", got)
		}
	})

	t.Run("menu without dynamic crumb keeps the Phase004 pair", func(t *testing.T) {
		got := render(true, false)
		if got != "      <Breadcrumb />\n      <Outlet />\n" {
			t.Errorf("got:\n%s", got)
		}
	})

	t.Run("no menu emits a bare Outlet", func(t *testing.T) {
		got := render(false, false)
		if got != "      <Outlet />\n" {
			t.Errorf("got:\n%s", got)
		}
	})
}

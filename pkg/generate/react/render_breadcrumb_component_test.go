//ff:func feature=gen-react type=test control=sequence
//ff:what renderBreadcrumbComponent — 라우트 매칭/1조각 trail 비렌더/Link·span 분기, dynamic 라벨 prop·정적 폴백 검증

package react

import "testing"

func TestRenderBreadcrumbComponent(t *testing.T) {
	t.Run("static (no data-crumb-field) keeps the Phase004 form", func(t *testing.T) {
		got := renderBreadcrumbComponent(false)
		assertContains(t, got, "import { Link, matchPath, useLocation } from 'react-router-dom'")
		assertContains(t, got, "import { BREADCRUMBS, BREADCRUMB_ROUTES } from '@/lib/breadcrumbs'")
		assertContains(t, got, "export function Breadcrumb()")
		assertContains(t, got, "BREADCRUMB_ROUTES.find((r) => matchPath(r.pattern, pathname))")
		// single-crumb trails (depth-1 pages) and unmatched routes render nothing
		assertContains(t, got, "if (!trail || trail.length < 2) return null")
		assertContains(t, got, "<nav aria-label=\"Breadcrumb\">")
		assertContains(t, got, "{crumb.href ? <Link to={crumb.href}>{crumb.label}</Link> : <span>{crumb.label}</span>}")
		assertNotContains(t, got, "label?")
	})

	t.Run("dynamic accepts the label prop and falls back to the static label", func(t *testing.T) {
		got := renderBreadcrumbComponent(true)
		assertContains(t, got, "export function Breadcrumb({ label }: { label?: string | null })")
		// the single fallback point: dynamic crumb + label state set → state, else static label
		assertContains(t, got, "<span>{crumb.dynamic && label != null ? label : crumb.label}</span>")
		assertContains(t, got, "if (!trail || trail.length < 2) return null")
	})
}

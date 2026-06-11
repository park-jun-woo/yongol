//ff:func feature=gen-react type=test control=sequence
//ff:what renderDynamicBreadcrumbComponent — label prop 시그니처·dynamic 분기 정적 폴백·비렌더 가드 소스 검증

package react

import "testing"

func TestRenderDynamicBreadcrumbComponent(t *testing.T) {
	got := renderDynamicBreadcrumbComponent()
	assertContains(t, got, "export function Breadcrumb({ label }: { label?: string | null })")
	assertContains(t, got, "<span>{crumb.dynamic && label != null ? label : crumb.label}</span>")
	assertContains(t, got, "if (!trail || trail.length < 2) return null")
	assertContains(t, got, "{crumb.href ? <Link to={crumb.href}>{crumb.label}</Link> :")
}

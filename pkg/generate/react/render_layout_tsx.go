//ff:func feature=gen-react type=generator control=sequence
//ff:what 단일 레이아웃 컴포넌트의 TSX 소스를 생성한다 (sitemap 메뉴/현행 data-nav 분기 + NavLink·lucide 아이콘 import 조건부 + ROLES 상수·userRole 클레임 셀렉터 + Outlet 위 Breadcrumb + 동적 crumb 라벨 state·pathname 리셋·Outlet context + 동적 메뉴 그룹 useQuery·bearer token 게이트 + data-logout 모드별 방출)

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderLayoutTSX generates the full TSX source for a single layout
// component. With a sitemap-derived menu (menu != nil — plans/stml/sitemap
// Phase003) the nav block comes from the sitemap tree (NavLink active
// states, lucide-react icon imports, useLocation/pathname for the
// ancestor highlight) and the shared <Breadcrumb> renders above the
// Outlet (plans/stml/sitemap Phase004 — it selects its trail from the
// current route, so layouts whose pages carry no trail render nothing).
// menu.DynamicCrumb (Phase006 — a hosted page declares data-crumb-field)
// adds the crumb-label state fed to <Breadcrumb label={...}>, a
// pathname-change reset (stale-label guard) and the setter handed down
// through <Outlet context={{ setCrumbLabel }}>; without it the
// Phase004/005 bytes stay identical. Dynamic menu groups (Phase007 —
// items with a data-fetch wiring) add one useQuery per distinct
// operation, keyed by the page fetch convention (['<Op>'] — so a page
// action's data-invalidates refreshes the menu too); bearer mode gates
// the queries on the stored token, cookie mode fires ungated. Without a
// sitemap (menu == nil) data-nav values resolve
// through navLinkPath byte-identically to the pre-sitemap output (static
// path verbatim, page-name reference → resolved route — page-flow
// Phase010). data-logout emits a mode-wired logout button when the
// project has auth (authMode "bearer" or "cookie" — the prepared.AuthFor
// derivation), and is skipped entirely without auth (TM-38 surfaces the
// dead declaration). The layout is the first component class importing
// the api client and session store — the import paths follow the page
// emitter convention (@/lib/api, @/stores/auth).
func renderLayoutTSX(componentName string, layout stml.LayoutSpec, routePatterns map[string]string, authMode string, menu *sitemapMenu) string {
	emitLogout := layout.Logout != nil && authMode != ""
	// dynamic crumb label (plans/stml/sitemap Phase006): a hosted page
	// declares data-crumb-field — the layout keeps the label state, resets
	// it on pathname change (stale-label guard) and hands the setter down
	// through the Outlet context. Needs the Outlet (no Outlet = no
	// Breadcrumb = nothing to label).
	dynamicCrumb := menu != nil && menu.DynamicCrumb && layout.HasOutlet
	usesLocation := menu != nil && (menuUsesLocation(menu.Items) || dynamicCrumb)
	// dynamic menu groups (plans/stml/sitemap Phase007): each distinct
	// data-fetch op gets one layout useQuery — query key = the page fetch
	// convention, so data-invalidates hits it. Bearer mode gates the
	// queries on the stored token (signed-out visitors never fire the
	// protected call); cookie mode fires ungated.
	var dynamicOps []string
	if menu != nil {
		dynamicOps = collectMenuDynamicOps(menu.Items)
	}
	hasDynamicGroups := len(dynamicOps) > 0
	bearerGate := hasDynamicGroups && authMode == "bearer"
	// data-roles menu filter (plans/stml/sitemap Phase005): active only
	// when the menu carries role-gated items AND role_field is wired —
	// otherwise the output stays byte-identical to the role-less emission.
	var roleSets [][]string
	if menu != nil && menu.RoleField != "" {
		roleSets = collectMenuRoleSets(menu.Items)
	}
	usesRoles := len(roleSets) > 0
	var sb strings.Builder

	imports := layoutImports(layout, emitLogout, menu)
	fmt.Fprintf(&sb, "import { %s } from 'react-router-dom'\n", strings.Join(imports, ", "))
	if dynamicCrumb {
		sb.WriteString("import { useEffect, useState } from 'react'\n")
	}
	if hasDynamicGroups {
		sb.WriteString("import { useQuery } from '@tanstack/react-query'\n")
	}
	if menu != nil {
		if icons := menuIconNames(menu.Items); len(icons) > 0 {
			fmt.Fprintf(&sb, "import { %s } from 'lucide-react'\n", strings.Join(icons, ", "))
		}
		if layout.HasOutlet {
			sb.WriteString("import { Breadcrumb } from '@/components/ui/Breadcrumb'\n")
		}
	}
	if (emitLogout && authMode == "bearer") || usesRoles || bearerGate {
		sb.WriteString("import { useAuthStore } from '@/stores/auth'\n")
	}
	if (emitLogout && layout.Logout.OperationID != "") || hasDynamicGroups {
		sb.WriteString("import { api } from '@/lib/api'\n")
	}
	if usesRoles {
		renderRoleConsts(&sb, roleSets)
	}

	fmt.Fprintf(&sb, "\nexport default function %s() {\n", componentName)
	if usesLocation {
		sb.WriteString("  const { pathname } = useLocation()\n")
	}
	if dynamicCrumb {
		sb.WriteString("  const [crumbLabel, setCrumbLabel] = useState<string | null>(null)\n")
		sb.WriteString("  useEffect(() => {\n    setCrumbLabel(null)\n  }, [pathname])\n")
	}
	if usesRoles {
		fmt.Fprintf(&sb, "  const userRole = useAuthStore((s) => s.claims[%s])\n", tsSingleQuote(menu.RoleField))
	}
	if bearerGate {
		sb.WriteString("  const token = useAuthStore((s) => s.token)\n")
	}
	renderSitemapGroupQueries(&sb, dynamicOps, bearerGate)
	if emitLogout {
		sb.WriteString("  const navigate = useNavigate()\n")
		sb.WriteString(renderLogoutHandler(layout.Logout.OperationID, authMode))
		sb.WriteString("\n")
	} else if usesLocation || usesRoles || dynamicCrumb || hasDynamicGroups {
		sb.WriteString("\n")
	}
	sb.WriteString("  return (\n")
	sb.WriteString("    <div>\n")

	renderLayoutNav(&sb, layout, routePatterns, emitLogout, menu)

	if layout.HasOutlet {
		renderLayoutOutlet(&sb, menu != nil, dynamicCrumb)
	}

	sb.WriteString("    </div>\n")
	sb.WriteString("  )\n")
	sb.WriteString("}\n")

	return sb.String()
}

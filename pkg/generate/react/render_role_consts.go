//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what renderRoleConsts — 레이아웃 모듈 레벨 ROLES_* 상수 선언들 방출

package react

import (
	"fmt"
	"strings"
)

// renderRoleConsts writes the module-level role allowlist constants of one
// layout (plans/stml/sitemap Phase005), one per distinct data-roles set:
//
//	const ROLES_admin_manager = ['admin', 'manager']
//
// The menu items' conditional renders reference them by the rolesConstName
// identifier.
func renderRoleConsts(sb *strings.Builder, roleSets [][]string) {
	sb.WriteString("\n")
	for _, set := range roleSets {
		quoted := make([]string, len(set))
		for i, r := range set {
			quoted[i] = tsSingleQuote(r)
		}
		fmt.Fprintf(sb, "const %s = [%s]\n", rolesConstName(set), strings.Join(quoted, ", "))
	}
}

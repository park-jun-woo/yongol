//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-46 — sitemap data-roles 값이 backend.auth.roles 에 없음 (ERROR; 메뉴 숨김 ≠ 접근 차단 명시)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm46SitemapRoleUnknown validates every sitemap data-roles value against
// manifest backend.auth.roles (plans/stml/sitemap Phase005): a typo'd role
// would silently never match any user's claim, hiding the menu entry from
// everyone. An empty backend.auth.roles list is TM-47's finding (broken
// wiring), so this rule stays silent then instead of flagging every value.
// The message states the industry separation explicitly: menu hiding is
// UX, not security — access blocking is Rego's (backend) concern, so a
// wrong role here never opens a hole, it only mis-renders the menu.
func tm46SitemapRoleUnknown(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	known := backendAuthRoles(fs)
	if len(known) == 0 {
		return nil
	}
	knownSet := make(map[string]bool, len(known))
	for _, r := range known {
		knownSet[r] = true
	}
	var diags []diagnostic.Diagnostic
	for _, e := range collectSitemapEntries(fs.Sitemap) {
		for _, role := range e.Node.Roles {
			if knownSet[role] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.Sitemap.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-46] sitemap data-roles value %q (at %s) is not in manifest backend.auth.roles — the menu entry would be hidden from every user (menu hiding is not security; access blocking is Rego's concern)", role, e.Path),
				Advice:  fmt.Sprintf("Use one of the declared roles (%s), or add %q to backend.auth.roles", strings.Join(known, ", "), role),
			})
		}
	}
	return diags
}

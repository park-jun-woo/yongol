//ff:func feature=gen-react type=accessor control=sequence
//ff:what resolveRoleField — manifest frontend.auth.role_field 접근자 (nil-safe)

package react

import "github.com/park-jun-woo/yongol/pkg/yongol"

// resolveRoleField returns manifest frontend.auth.role_field — the claim
// the sitemap data-roles menu filter reads (plans/stml/sitemap Phase005) —
// or "" when undeclared, which disables role conditions in the layout
// emission.
func resolveRoleField(fs *yongol.Fullstack) string {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Frontend.Auth == nil {
		return ""
	}
	return fs.Manifest.Frontend.Auth.RoleField
}

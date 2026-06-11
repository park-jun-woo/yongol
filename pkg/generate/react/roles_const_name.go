//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what rolesConstName — role 목록의 모듈 상수 식별자 (ROLES_<a>_<b>, 비식별자 문자는 _ 살균)

package react

import "strings"

// rolesConstName derives the module-level constant identifier holding one
// data-roles allowlist, e.g. ["admin","manager"] → "ROLES_admin_manager"
// (plans/stml/sitemap Phase005). Characters outside [A-Za-z0-9] are
// sanitized to '_' so a role like "super-admin" still yields a valid TS
// identifier; the constant value keeps the original strings.
func rolesConstName(roles []string) string {
	parts := make([]string, len(roles))
	for i, role := range roles {
		parts[i] = strings.Map(func(r rune) rune {
			switch {
			case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z', r >= '0' && r <= '9':
				return r
			default:
				return '_'
			}
		}, role)
	}
	return "ROLES_" + strings.Join(parts, "_")
}

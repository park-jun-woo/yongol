//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what backendAuthRoles — manifest backend.auth.roles 목록 반환 (nil-safe)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/yongol"

// backendAuthRoles returns manifest backend.auth.roles — the valid role
// name list the sitemap data-roles values must come from (TM-46) and whose
// emptiness breaks the data-roles wiring (TM-47). Nil when no backend.auth
// block (or roles list) is declared.
func backendAuthRoles(fs *yongol.Fullstack) []string {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	return fs.Manifest.Backend.Auth.Roles
}

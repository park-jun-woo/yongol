//ff:func feature=validate type=util control=selection topic=manifest-infra
//ff:what isAuthActive — manifest backend.auth 가 활성화되어 있는지 (type != "none")

package manifest_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// isAuthActive reports whether the manifest declares an active auth block.
// XDN-01~04 only fire when auth is active — `auth: nil` or
// `auth.type: none` (and the explicit empty string default for the missing
// block) skip the entire family.
func isAuthActive(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	auth := fs.Manifest.Backend.Auth
	if auth == nil {
		return false
	}
	switch auth.Type {
	case "none":
		return false
	default:
		return true
	}
}

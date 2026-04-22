//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what hasBearerAuth — manifest.backend.middleware 에 bearerAuth 포함 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasBearerAuth returns true when the manifest declares bearerAuth in
// backend.middleware.
func hasBearerAuth(fs *yongol.Fullstack) bool {
	if fs.Manifest == nil {
		return false
	}
	for _, mw := range fs.Manifest.Backend.Middleware {
		if mw == "bearerAuth" {
			return true
		}
	}
	return false
}

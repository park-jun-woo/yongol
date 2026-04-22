//ff:func feature=gen-gogin type=util control=sequence
//ff:what hasAuth — manifest.backend.auth 존재 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasAuth returns true when the manifest declares backend.auth config.
func hasAuth(fs *yongol.Fullstack) bool {
	return fs.Manifest != nil && fs.Manifest.Backend.Auth != nil && len(fs.Manifest.Backend.Auth.Claims) > 0
}

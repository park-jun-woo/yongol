//ff:func feature=generate type=util control=sequence
//ff:what manifestDeclaresSession — manifest.session.backend 선언 여부

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// manifestDeclaresSession returns true when the manifest explicitly
// declares session.backend. Nil-safe on every hop.
func manifestDeclaresSession(fs *yongol.Fullstack) bool {
	return fs != nil && fs.Manifest != nil &&
		fs.Manifest.Session != nil && fs.Manifest.Session.Backend != ""
}

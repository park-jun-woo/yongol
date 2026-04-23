//ff:func feature=generate type=util control=sequence
//ff:what manifestDeclaresCache — manifest.cache.backend 선언 여부

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// manifestDeclaresCache returns true when the manifest explicitly
// declares cache.backend. Nil-safe on every hop.
func manifestDeclaresCache(fs *yongol.Fullstack) bool {
	return fs != nil && fs.Manifest != nil &&
		fs.Manifest.Cache != nil && fs.Manifest.Cache.Backend != ""
}

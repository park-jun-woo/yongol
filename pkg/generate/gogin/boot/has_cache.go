//ff:func feature=gen-gogin type=util control=sequence
//ff:what hasCache — manifest.cache.backend 존재 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasCache returns true when the manifest declares cache.backend.
func hasCache(fs *yongol.Fullstack) bool {
	return fs.Manifest != nil && fs.Manifest.Cache != nil && fs.Manifest.Cache.Backend != ""
}

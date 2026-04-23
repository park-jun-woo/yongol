//ff:func feature=generate type=util control=sequence
//ff:what cacheBackendFor — manifest + SSaC 사용 여부로 cache 활성 판정 및 기본값 해석

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// cacheBackendFor returns non-nil iff the cache subsystem is in use.
// Mirrors sessionBackendFor. Default "memory" applied when SSaC uses
// cache but manifest is silent.
func cacheBackendFor(fs *yongol.Fullstack) *Cache {
	if manifestDeclaresCache(fs) {
		return &Cache{Backend: fs.Manifest.Cache.Backend}
	}
	if ssacUsesCacheCalls(fs) {
		return &Cache{Backend: "memory"}
	}
	return nil
}

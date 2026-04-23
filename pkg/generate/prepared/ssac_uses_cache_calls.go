//ff:func feature=generate type=util control=iteration dimension=2
//ff:what ssacUsesCacheCalls — SSaC 서비스 함수의 @call cache.* 사용 여부

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// ssacUsesCacheCalls mirrors pkg/validate/ssac_manifest.usesCache.
func ssacUsesCacheCalls(fs *yongol.Fullstack) bool {
	if fs == nil {
		return false
	}
	for _, fn := range fs.ServiceFuncs {
		if sequencesCallPrefix(fn.Sequences, "cache.") {
			return true
		}
	}
	return false
}

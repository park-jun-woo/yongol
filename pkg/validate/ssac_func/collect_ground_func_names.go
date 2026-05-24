//ff:func feature=validate type=util control=iteration dimension=1 topic=func-check
//ff:what collectGroundFuncNames — Ground Func.spec에서 지정 패키지의 함수 이름을 seen 맵에 추가

package ssac_func

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectGroundFuncNames adds function names from Ground's Func.spec lookup
// for the given package into the seen map.
func collectGroundFuncNames(pkg string, fs *yongol.Fullstack, seen map[string]bool) {
	g := fs.Ground()
	if g == nil {
		return
	}
	funcSpecs := g.Lookup["Func.spec"]
	if funcSpecs == nil {
		return
	}
	prefix := pkg + "."
	for key := range funcSpecs {
		if strings.HasPrefix(key, prefix) {
			camel := key[len(prefix):]
			seen[ucFirst(camel)] = true
		}
	}
}

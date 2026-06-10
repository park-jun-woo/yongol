//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what LinkRef의 파라미터 매핑을 세그먼트명 → 소스 맵으로 해석한다 (생략형은 유일한 필수 세그먼트에 귀속)
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// linkSegmentSources resolves a link's param bindings into a segment-name
// → source map against its TargetPattern. The elided form (empty Segment)
// binds to the single required segment of the target route; TM-32
// guarantees there is exactly one at validate time, so an ambiguous
// elision simply resolves to nothing here.
func linkSegmentSources(lr stmlparser.LinkRef) map[string]string {
	var required []string
	for _, seg := range strings.Split(lr.TargetPattern, "/") {
		if strings.HasPrefix(seg, ":") && !strings.HasSuffix(seg, "?") {
			required = append(required, strings.TrimPrefix(seg, ":"))
		}
	}
	out := map[string]string{}
	for _, p := range lr.Params {
		seg := p.Segment
		if seg == "" && len(required) == 1 {
			seg = required[0]
		}
		if seg != "" {
			out[seg] = p.Source
		}
	}
	return out
}

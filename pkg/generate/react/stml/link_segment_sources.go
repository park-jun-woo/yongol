//ff:func feature=stml-gen type=util control=sequence
//ff:what LinkRef의 파라미터 매핑을 세그먼트명 → 소스 맵으로 해석한다 (생략형은 유일한 필수 세그먼트에 귀속)
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// linkSegmentSources resolves a link's param bindings into a segment-name
// → source map against its TargetPattern. The elided form (empty Segment)
// binds to the single required segment of the target route; TM-32
// guarantees there is exactly one at validate time, so an ambiguous
// elision simply resolves to nothing here. The resolution core is shared
// with data-redirect emission (segmentSourcesFor).
func linkSegmentSources(lr stmlparser.LinkRef) map[string]string {
	return segmentSourcesFor(lr.TargetPattern, lr.Params)
}

//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 라우트 패턴과 파라미터 매핑을 세그먼트명 → 소스 맵으로 해석한다 (link/redirect 공유 코어, 생략형은 유일한 필수 세그먼트에 귀속)
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// segmentSourcesFor resolves "src -> Segment" bindings into a segment-name
// → source map against the given route pattern — the shared core of
// linkSegmentSources (data-link) and renderRedirectNavigate
// (data-redirect, page-flow Phase008). The elided form (empty Segment)
// binds to the single required segment of the pattern; TM-32/TM-33
// guarantee there is exactly one at validate time, so an ambiguous
// elision simply resolves to nothing here.
func segmentSourcesFor(pattern string, params []stmlparser.LinkParamBind) map[string]string {
	var required []string
	for _, seg := range strings.Split(pattern, "/") {
		if strings.HasPrefix(seg, ":") && !strings.HasSuffix(seg, "?") {
			required = append(required, strings.TrimPrefix(seg, ":"))
		}
	}
	out := map[string]string{}
	for _, p := range params {
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

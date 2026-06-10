//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what navLinkPath — data-nav 값 해석: "/" 시작은 정적 경로 그대로, 그 외 페이지명 → 해석 라우트(파라미터 세그먼트 strip)

package react

import "strings"

// navLinkPath resolves a layout data-nav value to the emitted <Link to>
// path (page-flow Phase010, the Phase008 dual rule): a "/"-prefixed value
// is a static path, emitted verbatim — byte-identical to the pre-Phase010
// output; any other value is a page-name reference resolved through
// routePatterns (falling back to "/<page-name>" when the page is unknown,
// like linkToAttr). Parameter segments are dropped defensively: optional
// segments (":Name?") simply mean "omitted", and required segments cannot
// survive validation (TM-36) — a menu link has no value to fill them.
func navLinkPath(target string, routePatterns map[string]string) string {
	if strings.HasPrefix(target, "/") {
		return target
	}
	pattern, ok := routePatterns[target]
	if !ok {
		return "/" + target
	}
	segs := strings.Split(pattern, "/")
	kept := segs[:0]
	for _, s := range segs {
		if strings.HasPrefix(s, ":") {
			continue
		}
		kept = append(kept, s)
	}
	out := strings.Join(kept, "/")
	if out == "" {
		return "/"
	}
	return out
}

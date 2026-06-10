//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what stripOptionalSegments — 라우트 패턴에서 ":Name?" optional 세그먼트를 제거한 base 경로 반환

package react

import "strings"

// stripOptionalSegments drops the trailing optional segments (":Name?",
// react-router v6.5+) from a route pattern and returns the base path a
// redirect can actually target — a <Navigate to> has no value to fill an
// optional segment, and an unfilled optional segment simply means "omitted"
// (page-flow Phase009). Required segments (":Name") are kept as-is; TM-34
// rejects them as redirect targets before generate runs.
func stripOptionalSegments(path string) string {
	segs := strings.Split(path, "/")
	kept := segs[:0]
	for _, s := range segs {
		if strings.HasPrefix(s, ":") && strings.HasSuffix(s, "?") {
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

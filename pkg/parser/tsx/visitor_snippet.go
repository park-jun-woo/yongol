//ff:func feature=tsx-parser type=accessor control=sequence
//ff:what (v *visitor).snippet — swc span 범위의 raw source bytes 반환

package tsx

import "strings"

// snippet returns the raw source bytes for a swc span (best-effort; used
// only for diagnostic values like ArgBinding.Value).
func (v *visitor) snippet(span astSpan) string {
	if span.Start <= 0 || span.End <= span.Start {
		return ""
	}
	s := span.Start - 1
	e := span.End - 1
	if s < 0 {
		s = 0
	}
	if e > len(v.src) {
		e = len(v.src)
	}
	if s >= e {
		return ""
	}
	return strings.TrimSpace(string(v.src[s:e]))
}

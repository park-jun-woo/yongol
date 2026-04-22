//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what warnDepthExceeds — MeasureDepth 결과에서 Q1 초과 라인을 WARN 문자열로 변환

package qcheck

import "fmt"

// warnDepthExceeds returns one WARN line per function whose max nesting
// depth exceeds lim.MaxDepth. The second return is the parser error from
// MeasureDepth if any; callers may decide to short-circuit on that.
func warnDepthExceeds(filename, src string, lim Limits) ([]string, error) {
	depths, err := MeasureDepth(filename, src)
	if err != nil {
		return nil, err
	}
	var warns []string
	for _, d := range depths {
		if d.MaxDepth <= lim.MaxDepth {
			continue
		}
		warns = append(warns, fmt.Sprintf("WARN: template emits depth=%d in %s func=%s — refactor (Q1 limit=%d)",
			d.MaxDepth, filename, d.Func, lim.MaxDepth))
	}
	return warns, nil
}

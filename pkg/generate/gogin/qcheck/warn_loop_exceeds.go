//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what warnLoopExceeds — MeasurePureLines 결과에서 Q4 초과 루프를 WARN 문자열로 변환

package qcheck

import "fmt"

// warnLoopExceeds returns one WARN line per for/range loop whose body pure
// line count exceeds lim.MaxPureLines. Parser errors bubble up so the caller
// can surface them separately from normal violations.
func warnLoopExceeds(filename, src string, lim Limits) ([]string, error) {
	loops, err := MeasurePureLines(filename, src)
	if err != nil {
		return nil, err
	}
	var warns []string
	for _, l := range loops {
		if l.PureLines <= lim.MaxPureLines {
			continue
		}
		warns = append(warns, fmt.Sprintf("WARN: template emits %s-body pure=%d in %s func=%s line=%d — refactor (Q4 limit=%d)",
			l.LoopKind, l.PureLines, filename, l.Func, l.Line, lim.MaxPureLines))
	}
	return warns, nil
}

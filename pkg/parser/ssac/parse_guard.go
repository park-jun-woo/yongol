//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @empty/@exists guard 시퀀스 파싱
package ssac

import (
	"strconv"
	"strings"
)

// parseGuard는 @empty/@exists를 파싱한다.
// target "message"
func parseGuard(seqType, rest string) *Sequence {
	rest = strings.TrimSpace(rest)
	target, msg, remainder := splitTargetMessage(rest)
	seq := &Sequence{
		Type:    seqType,
		Target:  target,
		Message: msg,
	}
	if remainder != "" {
		if code, err := strconv.Atoi(remainder); err == nil && code > 0 {
			seq.ErrStatus = code
		}
	}
	return seq
}

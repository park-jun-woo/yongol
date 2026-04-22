//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @call 외부 함수 호출 시퀀스 파싱
package ssac

import (
	"strconv"
	"strings"
)

// parseCall은 @call을 파싱한다.
// Type var = pkg.Func({Key: val, ...}) 또는 pkg.Func({Key: val, ...})
func parseCall(rest string) (*Sequence, error) {
	rest = strings.TrimSpace(rest)
	seq := &Sequence{Type: SeqCall}

	// = 가 있고, 그 전에 ( 가 없으면 result 있는 형태
	var remainder string
	eqIdx := strings.Index(rest, "=")
	parenIdx := strings.Index(rest, "(")
	if eqIdx > 0 && (parenIdx < 0 || eqIdx < parenIdx) {
		lhs := strings.TrimSpace(rest[:eqIdx])
		rhs := strings.TrimSpace(rest[eqIdx+1:])

		result := parseResult(lhs)
		if result == nil {
			return nil, nil
		}
		seq.Result = result

		model, inputs, rem, err := parseCallExprInputs(rhs)
		if err != nil {
			return nil, err
		}
		seq.Model = model
		seq.Inputs = inputs
		remainder = rem
	} else {
		model, inputs, rem, err := parseCallExprInputs(rest)
		if err != nil {
			return nil, err
		}
		seq.Model = model
		seq.Inputs = inputs
		remainder = rem
	}

	// trailing HTTP status code (e.g. "401")
	if remainder != "" {
		if code, err := strconv.Atoi(remainder); err == nil && code > 0 {
			seq.ErrStatus = code
		}
	}

	return seq, nil
}

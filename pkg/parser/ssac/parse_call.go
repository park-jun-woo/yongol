//ff:func feature=ssac-parse type=parser control=sequence
//ff:what parseCall — parses an @call external function call sequence
package ssac

import (
	"strconv"
	"strings"
)

// parseCall parses @call.
// Forms: Type var = pkg.Func({Key: val, ...}) or pkg.Func({Key: val, ...})
func parseCall(rest string) (*Sequence, error) {
	rest = strings.TrimSpace(rest)
	seq := &Sequence{Type: SeqCall}

	// If = is present and no ( appears before it, the form includes a result binding
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

//ff:func feature=ssac-parse type=parser control=sequence
//ff:what parseEval — parses an @eval predicate guard sequence (no result capture)
package ssac

import (
	"fmt"
	"strconv"
	"strings"
)

// parseEval parses @eval.
// Form: pkg.Func({Key: val, ...}) "message" STATUS
// Result capture (Type var =) is forbidden — @eval is a guard-only directive.
func parseEval(rest string) (*Sequence, error) {
	rest = strings.TrimSpace(rest)
	seq := &Sequence{Type: SeqEval}

	// Reject result-capture form: any "=" before "(" indicates "Type var = ...".
	eqIdx := strings.Index(rest, "=")
	parenIdx := strings.Index(rest, "(")
	if eqIdx > 0 && (parenIdx < 0 || eqIdx < parenIdx) {
		return nil, fmt.Errorf("@eval does not support result capture (Type var = ...); use @call for capturing return values")
	}

	model, inputs, remainder, err := parseCallExprInputs(rest)
	if err != nil {
		return nil, err
	}
	seq.Model = model
	seq.Inputs = inputs

	// remainder must be: "<message>" STATUS
	remainder = strings.TrimSpace(remainder)
	msg, rest2 := extractQuoted(remainder)
	seq.Message = msg
	rest2 = strings.TrimSpace(rest2)
	if rest2 != "" {
		if code, err := strconv.Atoi(rest2); err == nil && code > 0 {
			seq.ErrStatus = code
		}
	}
	return seq, nil
}

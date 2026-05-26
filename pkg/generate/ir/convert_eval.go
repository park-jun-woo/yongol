//ff:func feature=gen-ir type=util control=sequence
//ff:what convertEval -- @eval 시퀀스 → EvalOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertEval converts a @eval sequence to an IR Op.
func convertEval(seq ssac.Sequence) Op {
	pkg, fn := splitModelMethod(seq.Model)
	statusCode := seq.ErrStatus
	if statusCode == 0 {
		statusCode = 400
	}
	return Op{Kind: OpEval, Eval: &EvalOp{
		Package:    pkg,
		Function:   fn,
		Args:       convertInputsToFieldArgs(seq.Inputs),
		Message:    seq.Message,
		StatusCode: statusCode,
	}}
}

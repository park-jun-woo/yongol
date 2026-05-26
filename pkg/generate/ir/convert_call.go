//ff:func feature=gen-ir type=util control=sequence
//ff:what convertCall -- @call 시퀀스 → CallOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertCall converts a @call sequence to an IR Op.
func convertCall(seq ssac.Sequence) Op {
	pkg, fn := splitModelMethod(seq.Model)
	op := CallOp{
		Package:   pkg,
		Function:  fn,
		Args:      convertInputsToFieldArgs(seq.Inputs),
		ErrStatus: seq.ErrStatus,
		Message:   seq.Message,
	}
	if op.ErrStatus == 0 {
		op.ErrStatus = 500
	}
	if seq.Result != nil {
		op.ResultVar = seq.Result.Var
		op.ResultType = seq.Result.Type
	}
	return Op{Kind: OpCall, Call: &op}
}

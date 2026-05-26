//ff:func feature=gen-ir type=util control=sequence
//ff:what convertPost -- @post 시퀀스 → PostOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertPost converts a @post sequence to an IR Op.
func convertPost(seq ssac.Sequence) Op {
	model, method := splitModelMethod(seq.Model)
	op := PostOp{
		Model:  model,
		Method: method,
		Args:   convertInputsToFieldArgs(seq.Inputs),
	}
	if seq.Result != nil {
		op.VarName = seq.Result.Var
		op.VarType = seq.Result.Type
		op.IsList = seq.Result.Wrapper != ""
	}
	return Op{Kind: OpPost, Post: &op}
}

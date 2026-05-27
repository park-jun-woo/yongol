//ff:func feature=gen-ir type=util control=sequence
//ff:what convertGet -- @get 시퀀스 → GetOp IR 변환

package ir

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// convertGet converts a @get sequence to an IR Op.
func convertGet(seq ssac.Sequence) Op {
	model, method := splitModelMethod(seq.Model)
	op := GetOp{
		Model:  model,
		Method: method,
		Args:   convertInputsToFieldArgs(seq.Inputs),
	}
	if seq.Result != nil {
		op.VarName = seq.Result.Var
		op.VarType = seq.Result.Type
		op.IsList = seq.Result.Wrapper != "" || strings.HasPrefix(seq.Result.Type, "[]")
	}
	return Op{Kind: OpGet, Get: &op}
}

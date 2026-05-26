//ff:func feature=gen-ir type=util control=sequence
//ff:what convertDelete -- @delete 시퀀스 → DeleteOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertDelete converts a @delete sequence to an IR Op.
func convertDelete(seq ssac.Sequence) Op {
	model, method := splitModelMethod(seq.Model)
	return Op{Kind: OpDelete, Delete: &DeleteOp{
		Model:  model,
		Method: method,
		Args:   convertInputsToFieldArgs(seq.Inputs),
	}}
}

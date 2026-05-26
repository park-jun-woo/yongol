//ff:func feature=gen-ir type=util control=sequence
//ff:what convertPut -- @put 시퀀스 → PutOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertPut converts a @put sequence to an IR Op.
func convertPut(seq ssac.Sequence) Op {
	model, method := splitModelMethod(seq.Model)
	return Op{Kind: OpPut, Put: &PutOp{
		Model:  model,
		Method: method,
		Args:   convertInputsToFieldArgs(seq.Inputs),
	}}
}

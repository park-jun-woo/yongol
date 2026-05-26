//ff:func feature=gen-ir type=util control=sequence
//ff:what convertState -- @state 시퀀스 → StateOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertState converts a @state sequence to an IR Op.
func convertState(seq ssac.Sequence) Op {
	statusCode := seq.ErrStatus
	if statusCode == 0 {
		statusCode = 409
	}
	return Op{Kind: OpState, State: &StateOp{
		Diagram:    seq.DiagramID,
		Inputs:     convertInputsToFieldArgs(seq.Inputs),
		Transition: seq.Transition,
		Message:    seq.Message,
		StatusCode: statusCode,
	}}
}

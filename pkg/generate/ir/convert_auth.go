//ff:func feature=gen-ir type=util control=sequence
//ff:what convertAuth -- @auth 시퀀스 → AuthOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertAuth converts a @auth sequence to an IR Op.
func convertAuth(seq ssac.Sequence) Op {
	statusCode := seq.ErrStatus
	if statusCode == 0 {
		statusCode = 403
	}
	return Op{Kind: OpAuth, Auth: &AuthOp{
		Action:     seq.Action,
		Resource:   seq.Resource,
		Inputs:     convertInputsToFieldArgs(seq.Inputs),
		Message:    seq.Message,
		StatusCode: statusCode,
	}}
}

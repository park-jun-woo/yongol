//ff:func feature=gen-ir type=util control=sequence
//ff:what convertEmpty -- @empty 시퀀스 → EmptyOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertEmpty converts a @empty sequence to an IR Op.
func convertEmpty(seq ssac.Sequence) Op {
	statusCode := seq.ErrStatus
	if statusCode == 0 {
		statusCode = 404
	}
	return Op{Kind: OpEmpty, Empty: &EmptyOp{
		VarName:    seq.Target,
		Message:    seq.Message,
		StatusCode: statusCode,
	}}
}

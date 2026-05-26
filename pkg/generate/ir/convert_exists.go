//ff:func feature=gen-ir type=util control=sequence
//ff:what convertExists -- @exists 시퀀스 → ExistsOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertExists converts a @exists sequence to an IR Op.
func convertExists(seq ssac.Sequence) Op {
	statusCode := seq.ErrStatus
	if statusCode == 0 {
		statusCode = 409
	}
	return Op{Kind: OpExists, Exists: &ExistsOp{
		VarName:    seq.Target,
		Message:    seq.Message,
		StatusCode: statusCode,
	}}
}

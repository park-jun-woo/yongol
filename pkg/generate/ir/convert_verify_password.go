//ff:func feature=gen-ir type=util control=sequence
//ff:what convertVerifyPassword -- @verify-password 시퀀스 → VerifyPasswordOp IR 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// convertVerifyPassword converts a @verify-password sequence to an IR Op.
func convertVerifyPassword(seq ssac.Sequence) Op {
	op := VerifyPasswordOp{
		Model:        seq.Model,
		EmailCol:     seq.EmailCol,
		EmailExpr:    seq.EmailExpr,
		HashCol:      seq.HashCol,
		PasswordExpr: seq.PasswordExpr,
		ErrStatus:    seq.ErrStatus,
		Message:      seq.Message,
	}
	if seq.Result != nil {
		op.ResultVar = seq.Result.Var
		op.ResultType = seq.Result.Type
	}
	return Op{Kind: OpVerifyPassword, VerifyPW: &op}
}

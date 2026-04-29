//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what xqs20EligibleSeqType — XQS-20 적용 대상 SSaC 시퀀스 타입 화이트리스트

package ssac_sqlc

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// xqs20EligibleSeqType reports whether a sequence type carries a declared
// return-type that participates in the XQS-20 check. Only CRUD shapes that
// actually invoke sqlc and bind a result. `@delete` is excluded because it
// has no return, and `@call` / `@eval` are excluded because they target Go
// packages, not sqlc.
func xqs20EligibleSeqType(t string) bool {
	return t == ssac.SeqGet || t == ssac.SeqPost || t == ssac.SeqPut
}

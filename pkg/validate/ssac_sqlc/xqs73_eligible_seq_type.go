//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what xqs73EligibleSeqType — 시퀀스 타입이 sqlc 쿼리 결과를 생성하는지 판별

package ssac_sqlc

// xqs73EligibleSeqType returns true for sequence types that produce sqlc query results.
func xqs73EligibleSeqType(seqType string) bool {
	return seqType == "get" || seqType == "post" || seqType == "put"
}

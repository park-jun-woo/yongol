//ff:func feature=gen-ir type=util control=selection
//ff:what isCRUDSeq -- 시퀀스 타입이 sqlc 쿼리 대응 CRUD 인지 판정

package ir

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// isCRUDSeq returns true for sequence types that correspond to sqlc queries.
func isCRUDSeq(seqType string) bool {
	switch seqType {
	case ssac.SeqGet, ssac.SeqPost, ssac.SeqPut, ssac.SeqDelete:
		return true
	}
	return false
}

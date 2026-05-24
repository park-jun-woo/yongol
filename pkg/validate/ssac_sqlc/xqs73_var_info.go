//ff:type feature=validate type=rule topic=ssac-sqlc
//ff:what xqs73VarInfo — 부분 SELECT 쿼리와 결과 변수를 바인딩하는 중간 구조체

package ssac_sqlc

import (
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs73VarInfo holds the query and producing sequence for a result variable
// bound by a @get/@post/@put that uses a partial SELECT.
type xqs73VarInfo struct {
	query sqlcparser.QuerySpec
	seq   ssacparser.Sequence
}

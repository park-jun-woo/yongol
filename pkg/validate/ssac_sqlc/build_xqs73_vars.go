//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what buildXqs73Vars — SSaC 함수의 @get/@post/@put 결과에서 부분 SELECT 변수 맵 구축

package ssac_sqlc

import (
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildXqs73Vars builds a map from varName to (querySpec, producing sequence)
// for all @get/@post/@put results that reference partial SELECT queries.
func buildXqs73Vars(fn ssacparser.ServiceFunc, queryMap map[string]sqlcparser.QuerySpec) map[string]xqs73VarInfo {
	vars := make(map[string]xqs73VarInfo)
	for _, seq := range fn.Sequences {
		if !xqs73EligibleSeqType(seq.Type) {
			continue
		}
		if seq.Result == nil || seq.Result.Var == "" {
			continue
		}
		if seq.Package != "" {
			continue // @call — not a sqlc query
		}
		queryName := resolveQueryName(seq)
		q, ok := queryMap[queryName]
		if !ok {
			continue // XQS-19 covers missing query
		}
		// Only check partial SELECT queries
		if q.SelectStar || len(q.SelectCols) == 0 {
			continue
		}
		vars[seq.Result.Var] = xqs73VarInfo{query: q, seq: seq}
	}
	return vars
}

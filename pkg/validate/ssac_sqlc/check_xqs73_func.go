//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what checkXqs73Func — 하나의 SSaC 함수에서 부분 SELECT 필드 참조 위반 검사

package ssac_sqlc

import (
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// checkXqs73Func checks one SSaC function for partial SELECT field
// reference violations.
func checkXqs73Func(fn ssacparser.ServiceFunc, queryMap map[string]sqlcparser.QuerySpec) []diagnostic.Diagnostic {
	// Build a map: varName → (querySpec, producing sequence) for all @get/@post/@put results.
	vars := buildXqs73Vars(fn, queryMap)
	if len(vars) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, seq := range fn.Sequences {
		switch seq.Type {
		case "response":
			diags = append(diags, checkXqs73Response(fn, seq, vars)...)
		case "empty", "exists":
			diags = append(diags, checkXqs73Guard(fn, seq, vars)...)
		}
	}
	return diags
}

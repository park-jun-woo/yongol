//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what xqs74CheckFunc — 하나의 SSaC 함수에서 @empty/@exists guard의 non-integer PK 검사

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs74CheckFunc checks one SSaC function for @empty/@exists guards on
// non-integer PK models.
func xqs74CheckFunc(fn ssacparser.ServiceFunc, tableMap map[string]*ddl.Table) []diagnostic.Diagnostic {
	// Build map: varName → modelName from @get/@post result bindings.
	varModel := make(map[string]string)
	for _, seq := range fn.Sequences {
		if seq.Type != "get" && seq.Type != "post" {
			continue
		}
		if seq.Result == nil || seq.Result.Var == "" || seq.Result.Type == "" {
			continue
		}
		varModel[seq.Result.Var] = seq.Result.Type
	}
	if len(varModel) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	for _, seq := range fn.Sequences {
		if d, ok := xqs74CheckGuardSeq(seq, varModel, tableMap, fn.FileName); ok {
			diags = append(diags, d)
		}
	}
	return diags
}

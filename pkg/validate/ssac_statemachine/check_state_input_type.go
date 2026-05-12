//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-statemachine
//ff:what checkStateInputType — 단일 @state 시퀀스의 input value 타입이 string 호환인지 검증

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// checkStateInputType validates each input in a @state sequence. Returns
// diagnostics for inputs whose resolved Go type is not string-compatible.
func checkStateInputType(g *rule.Ground, fn parsessac.ServiceFunc, seq parsessac.Sequence) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for inputKey, inputValue := range seq.Inputs {
		sourceType := resolveStateInputType(g, fn.Name, inputValue)
		d := makeStateTypeDiag(fn.FileName, seq.Line, inputKey, sourceType)
		if d != nil {
			diags = append(diags, *d)
		}
	}
	return diags
}

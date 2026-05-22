//ff:func feature=validate type=rule control=iteration dimension=1 topic=func-check
//ff:what checkAuthInputType — 단일 @auth 시퀀스의 input value 타입이 string 호환인지 검증

package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// checkAuthInputType validates each input in an @auth sequence. Returns
// diagnostics for inputs whose resolved Go type is not string-compatible.
func checkAuthInputType(g *rule.Ground, fn parsessac.ServiceFunc, seq parsessac.Sequence) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for inputKey, inputValue := range seq.Inputs {
		sourceType := resolveInputType(g, fn.Name, inputValue)
		d := makeAuthTypeDiag(fn.FileName, seq.Line, inputKey, sourceType, fn.Name)
		if d != nil {
			diags = append(diags, *d)
		}
	}
	return diags
}

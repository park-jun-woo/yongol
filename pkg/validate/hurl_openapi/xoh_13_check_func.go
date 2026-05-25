//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh13CheckFunc — 단일 SSaC 함수의 guard + happy path hurl 커버리지 진단

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func xoh13CheckFunc(fn ssac.ServiceFunc, coveredSet map[string]bool) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, seq := range fn.Sequences {
		diags = append(diags, xoh13CheckGuard(fn, seq, coveredSet)...)
	}
	diags = append(diags, xoh13CheckHappyPath(fn, coveredSet)...)
	return diags
}

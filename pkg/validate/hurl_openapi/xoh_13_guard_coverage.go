//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-13 — SSaC guard ErrStatus + @response happy path 가 hurl 에서 테스트되는지 검증

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func xoh13GuardCoverage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil || len(fs.ServiceFuncs) == 0 || len(fs.HurlEntries) == 0 {
		return nil
	}

	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	covered := collectCoveredStatusCodes(fs.HurlEntries, routes)

	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		coveredSet := covered[fn.Name]
		diags = append(diags, xoh13CheckFunc(fn, coveredSet, fs.ParsedPolicies)...)
	}
	return diags
}

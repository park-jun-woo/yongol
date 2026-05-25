//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-12 — OpenAPI 선언 status code 중 hurl 미테스트 WARNING (5xx 제외)

package hurl_openapi

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func xoh12StatusCoverage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil || len(fs.HurlEntries) == 0 {
		return nil
	}

	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	covered := collectCoveredStatusCodes(fs.HurlEntries, routes)

	var diags []diagnostic.Diagnostic
	for _, route := range routes {
		d := xoh12CheckRoute(route, covered)
		diags = append(diags, d...)
	}

	sort.Slice(diags, func(i, j int) bool {
		return diags[i].Message < diags[j].Message
	})
	return diags
}

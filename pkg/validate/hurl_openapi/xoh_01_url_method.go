//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-01 — hurl 요청의 URL path + method 가 OpenAPI operation 으로 선언됨

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh01URLMethod enforces XOH-01: every hurl request must hit an
// OpenAPI operation that declares the same path *and* method.
//
// This rule folds together the former XOH-35 (path only) and XOH-36
// (method on matched path). Splitting them produced two diagnostics for
// a single typo; a unified judgement keeps the advice concise and lets
// the XOH-03/04/08 rules skip cleanly when there is no matched operation
// to compare against.
func xoh01URLMethod(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		if d := checkEntryURLMethod(e, routes); d != nil {
			diags = append(diags, *d)
		}
	}
	return diags
}

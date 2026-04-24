//ff:func feature=validate type=rule control=sequence topic=query-rego
//ff:what Run — execute all sqlc↔Rego cross-validation rules (XQP-*)
package query_rego

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all sqlc ↔ Rego cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xqp30OwnerLookupQuery(fs)...)
	return diags
}

//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Run — sqlc 쿼리 standalone 검증 (Q-*)

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all sqlc query validation rules (Q-*). Covers standalone query
// sanity: cardinality, naming, WHERE presence, sensitive column leaks, etc.
// Cross-SSOT checks (query ↔ DDL / query ↔ OpenAPI) belong to PhaseV07.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, q01NameRequired(fs)...)
	diags = append(diags, q02Cardinality(fs)...)
	diags = append(diags, q03NamePascalCase(fs)...)
	diags = append(diags, q04ManyLimit(fs)...)
	diags = append(diags, q05DeleteWhere(fs)...)
	diags = append(diags, q06UpdateWhere(fs)...)
	diags = append(diags, q07SelectStarSensitive(fs)...)
	diags = append(diags, q08UnusedParam(fs)...)
	diags = append(diags, q09SelectOnExec(fs)...)
	diags = append(diags, q11SqlPackagePgxV5(fs)...)
	return diags
}

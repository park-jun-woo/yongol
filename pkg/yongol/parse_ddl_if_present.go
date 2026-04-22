//ff:func feature=orchestrator type=loader control=sequence
//ff:what DDL 탐지 시 results/tables/sqlc queries 3종을 모두 파싱
package yongol

import (
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
)

// parseDDLIfPresent loads DDL results, tables and sqlc queries when DDL is present.
// sqlc queries live under <ddl>/queries/ (e.g. specs/db/queries/).
func parseDDLIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindDDL]
	if !ok {
		return
	}
	results, diags := ddl.ParseDir(d.Path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.DDLResults = results
	}
	tables, tdiags := ddl.ParseTables(d.Path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, tdiags...)
	if len(tdiags) == 0 {
		fs.DDLTables = tables
	}
	queries, qdiags := sqlcparser.ParseDir(filepath.Join(d.Path, "queries"))
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, qdiags...)
	if len(qdiags) == 0 {
		fs.SQLcQueries = queries
	}
}

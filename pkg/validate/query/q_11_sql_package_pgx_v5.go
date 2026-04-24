//ff:func feature=validate type=rule control=iteration dimension=1 topic=query-structural
//ff:what Q-11 — sqlc.yaml sql_package 이 pgx/v5 가 아니면 ERROR

package query

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q11SqlPackagePgxV5 validates Q-11: every sql[].gen.go.sql_package in
// db/sqlc.yaml must be "pgx/v5". yongol's backend codegen (handler tx,
// convert pgtype unwrap, server bootstrap) is unified on pgx/v5 — any other
// value (database/sql, pgx/v4, absent, lib/pq, ...) makes the generated
// backend fail to compile. Rejecting at validate time surfaces the fix one
// step earlier than a `go build` breakage after generate.
//
// The rule reads sqlc.yaml directly rather than going through the existing
// ddl-level helpers because it needs the `sql_package` field which
// parseSqlcYaml (pkg/validate/ddl) intentionally omits from its minimal
// subset.
func q11SqlPackagePgxV5(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.SpecsDir == "" {
		return nil
	}
	sqlcPath := filepath.Join(fs.SpecsDir, "db", "sqlc.yaml")
	data, err := os.ReadFile(sqlcPath)
	if err != nil {
		// D-4 already reports sqlc.yaml missing; stay silent here.
		return nil
	}
	var cfg sqlcPackageConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for i, entry := range cfg.SQL {
		if d := diagnoseSqlcPackageEntry(i, entry.Gen.Go.SqlPackage); d != nil {
			diags = append(diags, *d)
		}
	}
	return diags
}

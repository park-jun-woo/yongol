//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Q-11 — sqlc.yaml sql_package 이 pgx/v5 가 아니면 ERROR

package query

import (
	"fmt"
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
	var cfg struct {
		SQL []struct {
			Gen struct {
				Go struct {
					SqlPackage string `yaml:"sql_package"`
				} `yaml:"go"`
			} `yaml:"gen"`
		} `yaml:"sql"`
	}
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for i, entry := range cfg.SQL {
		pkg := entry.Gen.Go.SqlPackage
		if pkg == "pgx/v5" {
			continue
		}
		current := pkg
		if current == "" {
			current = "(absent; sqlc defaults to database/sql)"
		} else {
			current = fmt.Sprintf("%q", current)
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:  "db/sqlc.yaml",
			Line:  0,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: fmt.Sprintf(
				"[Q-11] sqlc.yaml sql[%d].gen.go.sql_package must be \"pgx/v5\" (current: %s)",
				i, current,
			),
			Advice: "yongol's backend codegen is unified on pgx/v5. Update db/sqlc.yaml:\n" +
				"  gen:\n" +
				"    go:\n" +
				"      sql_package: pgx/v5\n" +
				"Then re-run `yongol generate <specs> <arts>`. database/sql / pgx/v4 support was removed.",
		})
	}
	return diags
}

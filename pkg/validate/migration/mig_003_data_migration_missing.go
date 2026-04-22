//ff:func feature=validate type=rule control=iteration dimension=1 topic=migration-hints
//ff:what MIG-003 — @data_migration file=... 경로에 파일 없음 → ERROR
package migration

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// Mig003DataMigrationMissing takes the list of missing sidecar paths
// (returned by migration.LoadDataMigrationSQL) and converts them into
// diagnostics.
func Mig003DataMigrationMissing(missing []string) []diagnostic.Diagnostic {
	var out []diagnostic.Diagnostic
	for _, p := range missing {
		out = append(out, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[MIG-003] @data_migration file not found: " + p,
			Advice:  "create the sidecar file or fix the path in the hint",
		})
	}
	return out
}

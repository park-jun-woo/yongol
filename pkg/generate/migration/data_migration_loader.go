//ff:func feature=migration type=parser control=sequence
//ff:what DataMigrationLoader — @data_migration file=... sidecar SQL 읽기 + 누락 검증
package migration

import (
	"os"
	"path/filepath"
)

// LoadDataMigrationSQL resolves every @data_migration sidecar referenced
// by hints and returns a map table -> SQL body. Missing files surface
// via the second return slice (MIG-003 input).
//
// specsDir is where DDL SSOTs live; sidecar paths in the hint are
// resolved relative to specsDir.
func LoadDataMigrationSQL(specsDir string, hints *Hints) (map[string]string, []string) {
	if hints == nil || len(hints.DataMigrations) == 0 {
		return nil, nil
	}
	out := map[string]string{}
	var missing []string
	for table, rel := range hints.DataMigrations {
		abs := rel
		if !filepath.IsAbs(rel) {
			abs = filepath.Join(specsDir, rel)
		}
		data, err := os.ReadFile(abs)
		if err != nil {
			missing = append(missing, rel)
			continue
		}
		out[table] = string(data)
	}
	return out, missing
}

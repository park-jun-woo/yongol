//ff:func feature=validate type=rule control=sequence topic=ddl-structural
//ff:what D-4 — db/sqlc.yaml 미존재 시 ERROR

package ddl

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// d04SqlcYamlRequired validates D-4: when the db/ directory is present (DDL
// detected), db/sqlc.yaml must also exist. sqlc.yaml is the user-owned config
// that yongol generate uses to invoke `sqlc generate`.
func d04SqlcYamlRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.PresenceOf(yongol.KindDDL) == yongol.SSOTAbsent {
		return nil
	}
	dbDir := filepath.Join(fs.SpecsDir, "db")
	sqlcPath := filepath.Join(dbDir, "sqlc.yaml")
	if _, err := os.Stat(sqlcPath); err == nil {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "db/sqlc.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[D-4] db/sqlc.yaml not found — yongol generate requires sqlc config",
		Advice:  "db/sqlc.yaml 을 작성하세요 (sqlc v2 format)",
	}}
}

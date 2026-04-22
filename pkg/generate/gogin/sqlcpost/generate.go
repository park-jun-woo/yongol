//ff:func feature=gen-gogin type=command control=iteration dimension=1
//ff:what Generate — sqlc row struct 에 LogValue 메소드 후처리 주입

package sqlcpost

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate emits one `<table>_log.go` file per DDL table that contains at
// least one `-- @sensitive` column. Each file defines a `LogValue()` method
// on the sqlc-generated row struct that returns a slog.GroupValue with
// sensitive fields replaced by "[REDACTED]". slog's LogValuer interface
// picks it up automatically when a handler logs the struct whole.
//
// sqlc itself is an external tool — yongol does not re-emit models.go.
// These sibling files live in the same db package so the methods attach
// to the sqlc row types.
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	if len(fs.DDLTables) == 0 {
		return nil
	}
	dbDir := filepath.Join(artifactsDir, "backend", "internal", "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		return err
	}
	for _, t := range fs.DDLTables {
		if len(t.SensitiveColumns) == 0 {
			continue
		}
		body, err := renderLogValueFile(t)
		if err != nil {
			return fmt.Errorf("table %s: %w", t.Name, err)
		}
		path := filepath.Join(dbDir, t.Name+"_log.go")
		if err := os.WriteFile(path, []byte(body), 0o644); err != nil {
			return fmt.Errorf("write %s: %w", path, err)
		}
	}
	return nil
}

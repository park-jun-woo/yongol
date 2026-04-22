//ff:func feature=gen-gogin type=util control=sequence
//ff:what generateSQLCGoGin — sqlc generate 실행 (db/sqlc.yaml 기반)

package gogin

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// generateSQLCGoGin invokes `sqlc generate` using db/sqlc.yaml. The sqlc.yaml
// file is user-authored and specifies schema (DDL), queries, engine, and Go
// output directory. yongol does not generate or modify sqlc.yaml — it only
// runs sqlc with it.
//
// The working directory is set to db/ so that sqlc.yaml's relative paths
// (schema: ".", queries: "queries/") resolve correctly.
func generateSQLCGoGin(specsDir string) error {
	dbDir := filepath.Join(specsDir, "db")

	cmd := exec.Command("sqlc", "generate", "-f", "sqlc.yaml")
	cmd.Dir = dbDir
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("sqlc generate: %w — %s", err, string(out))
	}
	return nil
}

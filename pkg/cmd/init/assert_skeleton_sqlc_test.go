//ff:func feature=cli-init type=test-helper control=sequence
//ff:what assertSkeletonSqlc — sqlc.yaml contains pgx/v5 sql_package directive

package cliinit

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func assertSkeletonSqlc(t *testing.T, target string) {
	t.Helper()
	sqlc, err := os.ReadFile(filepath.Join(target, "specs/db/sqlc.yaml"))
	if err != nil {
		t.Fatalf("read sqlc.yaml: %v", err)
	}
	if !strings.Contains(string(sqlc), "sql_package: pgx/v5") {
		t.Errorf("sqlc.yaml missing pgx/v5 directive")
	}
}

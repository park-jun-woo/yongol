//ff:func feature=cli type=util control=sequence
//ff:what printMigrationLatestFile — artifacts migrations 디렉토리에서 가장 최근 .sql 파일명 출력
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// printMigrationLatestFile writes the "latest: <file>" line when an
// artifacts/db/migrations/*.sql file is present. A missing dir is silently
// ignored — this line is purely informational.
func printMigrationLatestFile(w io.Writer, artsDir string) {
	if artsDir == "" {
		return
	}
	migDir := filepath.Join(artsDir, migration.MigrationsSubdir)
	entries, err := os.ReadDir(migDir)
	if err != nil {
		return
	}
	latest := latestSQLFileName(entries)
	if latest == "" {
		return
	}
	fmt.Fprintf(w, "latest:   %s\n", filepath.Join(migration.MigrationsSubdir, latest))
}

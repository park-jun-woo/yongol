//ff:func feature=cli type=util control=iteration dimension=1
//ff:what printMigrationPendingDiff — prev↔curr 스키마 diff 개수 + 최대 10건 요약 출력
package main

import (
	"fmt"
	"io"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// printMigrationPendingDiff writes the "pending: N change(s)" summary and up
// to 10 change descriptions. Returns false on parse error so the caller stops
// the dashboard section.
func printMigrationPendingDiff(w io.Writer, ddlDir, snapshotPath string) bool {
	prev, _ := migration.LoadSnapshot(snapshotPath)
	if prev == nil {
		prev = migration.NewSchema()
	}
	curr, err := migration.BuildASTFromDir(ddlDir, nil)
	if err != nil {
		fmt.Fprintf(w, "pending: (parse error: %s)\n", err)
		return false
	}
	rawHints, _ := ddl.ExtractHintCommentsFromDir(ddlDir)
	hints := migration.ParseHints(rawHints)
	ops := migration.Diff(prev, curr, hints)
	fmt.Fprintf(w, "pending: %d change(s)\n", len(ops))
	max := len(ops)
	if max > 10 {
		max = 10
	}
	for i := 0; i < max; i++ {
		fmt.Fprintf(w, "  * %s\n", ops[i].Description())
	}
	if len(ops) > 10 {
		fmt.Fprintf(w, "  ... and %d more\n", len(ops)-10)
	}
	return true
}

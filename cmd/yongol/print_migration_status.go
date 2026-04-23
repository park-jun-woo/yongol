//ff:func feature=cli type=util control=iteration dimension=2
//ff:what printMigrationStatus — prints the Migration section of the yongol status dashboard
package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/migration"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// printMigrationStatus writes a Migration Status section to w. It is
// defensive: any error is turned into a single status line so the
// dashboard never fails.
func printMigrationStatus(w io.Writer, specsDir, artsDir string) {
	fmt.Fprintln(w, "\n── Migration Status ──")
	ddlDir := filepath.Join(specsDir, migration.SnapshotSubdir)
	snapshotPath := filepath.Join(ddlDir, migration.SnapshotFileName)

	// Snapshot info
	data, err := os.ReadFile(snapshotPath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintf(w, "snapshot: (absent — next generate will emit 0001_initial.sql)\n")
		} else {
			fmt.Fprintf(w, "snapshot: ERROR %s\n", err)
			return
		}
	} else {
		hash, drift := snapshotHashInfo(data)
		short := hash
		if len(short) > 8 {
			short = short[:8]
		}
		mark := "ok"
		if drift {
			mark = "DRIFT (MIG-006)"
		}
		fmt.Fprintf(w, "snapshot: %s  hash=%s  %s\n",
			filepath.Join(migration.SnapshotSubdir, migration.SnapshotFileName), short, mark)
	}

	// Pending diff — reload prev + curr and count ops.
	prev, _ := migration.LoadSnapshot(snapshotPath)
	if prev == nil {
		prev = migration.NewSchema()
	}
	curr, err := migration.BuildASTFromDir(ddlDir, []string{migration.SnapshotFileName})
	if err != nil {
		fmt.Fprintf(w, "pending: (parse error: %s)\n", err)
		return
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
	// Latest migration file.
	if artsDir != "" {
		migDir := filepath.Join(artsDir, migration.MigrationsSubdir)
		if entries, err := os.ReadDir(migDir); err == nil {
			latest := ""
			for _, e := range entries {
				name := e.Name()
				if !strings.HasSuffix(name, ".sql") {
					continue
				}
				if name > latest {
					latest = name
				}
			}
			if latest != "" {
				fmt.Fprintf(w, "latest:   %s\n", filepath.Join(migration.MigrationsSubdir, latest))
			}
		}
	}
}


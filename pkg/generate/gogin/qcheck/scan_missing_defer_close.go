//ff:func feature=gen-gogin type=util control=sequence
//ff:what ScanMissingDeferClose — DF-06 `db.Query` / `os.Open` 후 `defer <v>.Close()` 누락 탐지

package qcheck

import (
	"go/parser"
	"go/token"
)

// ScanMissingDeferClose parses src and returns one DefensiveFinding per
// resource-returning call (db.Query, db.QueryContext, os.Open) whose first
// LHS variable is not closed via `defer <v>.Close()` anywhere later in
// the same block. sqlc-generated *.sql.go files are the repository's main
// exception — callers should filter those out before passing src.
// Resource call targets are kept small on purpose; categories B/C expand
// in later Phases.
func ScanMissingDeferClose(filename, src string) ([]DefensiveFinding, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	return walkResourceBlocks(file, fset), nil
}

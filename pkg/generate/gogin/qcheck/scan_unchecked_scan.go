//ff:func feature=gen-gogin type=util control=sequence
//ff:what ScanUncheckedScan — DF-02 `row.Scan` / `rows.Scan` 에러 무시 탐지

package qcheck

import (
	"go/parser"
	"go/token"
)

// ScanUncheckedScan parses src and returns one DefensiveFinding per
// `<ident>.Scan(...)` call whose error result is not guarded. Both the
// pgx/v5 (pgx.Rows / pgx.Row) and the legacy database/sql Scan APIs
// return an error, so the same guard shapes as DF-01 apply: IfStmt.Init,
// two-line assign+guard, or the rarer `err = row.Scan(...)` followed by
// guard. Receiver ident is accepted as any name (not just row/rows)
// because sqlc variants like `r.Scan` exist in hand-authored wrappers.
func ScanUncheckedScan(filename, src string) ([]DefensiveFinding, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	return walkScanBlocks(file, fset), nil
}

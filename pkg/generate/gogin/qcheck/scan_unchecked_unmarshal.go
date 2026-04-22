//ff:func feature=gen-gogin type=util control=sequence
//ff:what ScanUncheckedUnmarshal — DF-01 `json.Unmarshal` / `yaml.Unmarshal` 에러 무시 탐지

package qcheck

import (
	"go/parser"
	"go/token"
)

// ScanUncheckedUnmarshal parses src and returns one DefensiveFinding per
// json.Unmarshal / yaml.Unmarshal call whose error return is not guarded.
// A call is considered guarded when (a) it lives inside an IfStmt.Init as
// the RHS of `err := json.Unmarshal(...)` followed by a nil-check, or
// (b) it is the RHS of a top-level AssignStmt whose LHS includes an
// identifier reused by the immediately following IfStmt's err-guard.
// All other shapes — bare expression statement, `_ = json.Unmarshal(...)`
// — are reported. Parser errors propagate unchanged.
func ScanUncheckedUnmarshal(filename, src string) ([]DefensiveFinding, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.SkipObjectResolution)
	if err != nil {
		return nil, err
	}
	return walkUnmarshalBlocks(file, fset), nil
}

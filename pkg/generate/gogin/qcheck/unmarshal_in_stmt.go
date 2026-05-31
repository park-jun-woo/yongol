//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what unmarshalInStmt — 단일 Stmt 에서 targets 의 Unmarshal 호출과 가드 여부 판정

package qcheck

import (
	"go/ast"
	"go/token"
)

// unmarshalInStmt inspects stmt (located at blockList[i]) and returns a
// DefensiveFinding for each targets.Unmarshal call whose error is not
// guarded. Three shapes are tolerated:
//   - `if err := pkg.Unmarshal(...); err != nil { ... }`
//   - `err := pkg.Unmarshal(...)` followed by an IfStmt err-guard
//   - `<x>, err := pkg.Unmarshal(...)` same shape
//
// Anything else (ExprStmt, `_ = pkg.Unmarshal(...)`, followed by a
// non-guard statement) is a finding.
func unmarshalInStmt(stmt ast.Stmt, blockList []ast.Stmt, i int, targets []string, fset *token.FileSet) []DefensiveFinding {
	var findings []DefensiveFinding
	for _, pkg := range targets {
		if ifStmt, ok := stmt.(*ast.IfStmt); ok && ifInitCalls(ifStmt, pkg, "Unmarshal") {
			continue
		}
		assign, ok := stmt.(*ast.AssignStmt)
		if ok && assignCallsAndGuarded(assign, pkg, "Unmarshal", blockList, i) {
			continue
		}
		if call := findCallInStmt(stmt, pkg, "Unmarshal"); call != nil {
			findings = append(findings, DefensiveFinding{
				Category: "DF-01",
				Detail:   pkg + ".Unmarshal",
				Pos:      fset.Position(call.Pos()),
			})
		}
	}
	return findings
}

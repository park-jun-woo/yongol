//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what BuildASTFromSQL — SQL 텍스트의 각 statement 를 Schema AST 에 누적
package migration

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// BuildASTFromSQL appends the structures parsed from one SQL text into
// the given Schema. It is the main lexer/parser entry point and is
// exported so tests can feed SQL literals directly.
func BuildASTFromSQL(s *Schema, sqlText string) error {
	// First, collect `-- @sentinel` INSERT blocks from the raw text.
	// The line-comment stripper used by splitStatements would otherwise
	// destroy the annotation marker.
	for _, r := range ddl.ScanSentinelInserts(sqlText) {
		if !r.Annotated {
			continue
		}
		name := canonIdent(r.Table)
		t := ensureTable(s, name)
		t.Sentinels = append(t.Sentinels, SentinelInsert{SQL: r.SQL})
	}

	stmts := splitStatements(sqlText)
	for _, stmt := range stmts {
		stmt = strings.TrimSpace(stmt)
		if err := dispatchStatement(s, stmt); err != nil {
			return err
		}
	}
	return nil
}

//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what BuildASTFromSQL — SQL 텍스트의 각 statement 를 Schema AST 에 누적
package migration

import "strings"

// BuildASTFromSQL appends the structures parsed from one SQL text into
// the given Schema. It is the main lexer/parser entry point and is
// exported so tests can feed SQL literals directly.
func BuildASTFromSQL(s *Schema, sqlText string) error {
	stmts := splitStatements(sqlText)
	for _, stmt := range stmts {
		stmt = strings.TrimSpace(stmt)
		if err := dispatchStatement(s, stmt); err != nil {
			return err
		}
	}
	return nil
}

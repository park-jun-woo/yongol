//ff:func feature=migration type=test control=sequence
//ff:what TestSplitStateFinish — splitState.finish trailing statement 처리 커버
package migration

import "testing"

func TestSplitStateFinishMethod(t *testing.T) {
	// no trailing semicolon → finish() must flush the last statement.
	stmts := splitStatements("SELECT 1; SELECT 2 /* x */ FROM t")
	if len(stmts) != 2 {
		t.Errorf("expected 2 statements, got %d: %v", len(stmts), stmts)
	}
}

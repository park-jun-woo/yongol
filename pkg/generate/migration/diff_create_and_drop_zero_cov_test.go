//ff:func feature=migration type=test control=sequence
//ff:what TestMigrationE2EZeroCov — ParseHints / BuildASTFromSQL / Diff / ApplyHintsToOps / EmitSQL 풀 파이프라인 커버
package migration

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestDiffCreateAndDrop_ZeroCov(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE gone (id BIGINT PRIMARY KEY);`)
	curr := mustAST(t, `CREATE TABLE fresh (id BIGINT PRIMARY KEY, label TEXT NOT NULL);`)
	hints := ParseHints([]ddl.HintComment{{Tag: "allow_destructive", TableCtx: "gone"}})
	ops := Diff(prev, curr, hints)
	withHints := ApplyHintsToOps(ops, hints)
	sql := EmitSQL(withHints, EmitOptions{})
	if !strings.Contains(sql, "CREATE TABLE fresh") {
		t.Errorf("missing CREATE TABLE fresh:\n%s", sql)
	}
}

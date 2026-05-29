//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XSD-55 — 미참조 DDL 테이블 검출 + @func-managed/@archived 면제

package ssac_ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// newXsd55Fullstack builds a Fullstack with the given DDL tables (no SSaC
// references) and a Ground carrying the provided flags, so xsd55DDLToModelRef
// can consult exemption flags.
func newXsd55Fullstack(flags rule.StringSet, tables ...ddl.Table) *yongol.Fullstack {
	fs := &yongol.Fullstack{DDLTables: tables}
	g := &rule.Ground{Flags: flags}
	fs.SetGround(g)
	return fs
}

func TestXsd55DDLToModelRef(t *testing.T) {
	// 1. Unreferenced, no annotation → ERROR (regression guard).
	t.Run("orphan table errors", func(t *testing.T) {
		fs := newXsd55Fullstack(rule.StringSet{}, ddl.Table{Name: "orphans"})
		diags := xsd55DDLToModelRef(fs)
		if len(diags) != 1 {
			t.Fatalf("want 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "orphans") {
			t.Errorf("diag message missing table name: %q", diags[0].Message)
		}
	})

	// 2. @func-managed table → exempt (no ERROR).
	t.Run("func-managed exempt", func(t *testing.T) {
		fs := newXsd55Fullstack(
			rule.StringSet{"funcManaged.bids": true},
			ddl.Table{Name: "bids"},
		)
		if diags := xsd55DDLToModelRef(fs); len(diags) != 0 {
			t.Fatalf("want 0 diags for func-managed table, got %d: %v", len(diags), diags)
		}
	})

	// 3. @archived table → still exempt (existing behavior preserved).
	t.Run("archived exempt", func(t *testing.T) {
		fs := newXsd55Fullstack(
			rule.StringSet{"archived.legacy": true},
			ddl.Table{Name: "legacy"},
		)
		if diags := xsd55DDLToModelRef(fs); len(diags) != 0 {
			t.Fatalf("want 0 diags for archived table, got %d: %v", len(diags), diags)
		}
	})

	// 4. Only the func-managed table is exempt; a sibling orphan still errors.
	t.Run("func-managed does not exempt unrelated orphan", func(t *testing.T) {
		fs := newXsd55Fullstack(
			rule.StringSet{"funcManaged.bids": true},
			ddl.Table{Name: "bids"},
			ddl.Table{Name: "orphans"},
		)
		diags := xsd55DDLToModelRef(fs)
		if len(diags) != 1 {
			t.Fatalf("want 1 diag, got %d: %v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "orphans") {
			t.Errorf("expected orphans to error, got: %q", diags[0].Message)
		}
	})
}

//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS74Branches_ZeroCov — xqs74CheckFunc/GuardSeq/Model 잔여 분기 직접 커버
package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXQS74CheckModel_Skips_ZeroCov(t *testing.T) {
	seq := ssacparser.Sequence{Type: "empty", Line: 3}

	// table not found
	if _, ok := xqs74CheckModel(seq, "Ghost", map[string]*ddl.Table{}, "f"); ok {
		t.Error("missing table should be skipped")
	}
	// table with no PK
	noPK := map[string]*ddl.Table{"users": {Name: "users", Columns: map[string]ddl.Column{}}}
	if _, ok := xqs74CheckModel(seq, "User", noPK, "f"); ok {
		t.Error("table with no PK should be skipped")
	}
	// PK column missing from Columns map
	pkMissingCol := map[string]*ddl.Table{"users": {Name: "users", PrimaryKey: []string{"id"}, Columns: map[string]ddl.Column{}}}
	if _, ok := xqs74CheckModel(seq, "User", pkMissingCol, "f"); ok {
		t.Error("PK col missing from Columns should be skipped")
	}
	// integer PK → no violation
	intPK := map[string]*ddl.Table{"users": {Name: "users", PrimaryKey: []string{"id"}, Columns: map[string]ddl.Column{"id": {Name: "id", RawType: "BIGINT"}}}}
	if _, ok := xqs74CheckModel(seq, "User", intPK, "f"); ok {
		t.Error("integer PK should not be a violation")
	}
	// non-integer PK → violation
	textPK := map[string]*ddl.Table{"users": {Name: "users", PrimaryKey: []string{"token"}, Columns: map[string]ddl.Column{"token": {Name: "token", RawType: "TEXT"}}}}
	d, ok := xqs74CheckModel(seq, "User", textPK, "f")
	if !ok {
		t.Fatal("text PK should be a violation")
	}
	if d.Line != 3 || d.File != "f" {
		t.Errorf("diag loc = %q:%d", d.File, d.Line)
	}
}

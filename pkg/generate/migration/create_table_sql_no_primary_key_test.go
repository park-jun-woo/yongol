//ff:func feature=migration type=test control=sequence
//ff:what TestCreateTable_SQL — 컬럼 + PRIMARY KEY 절 렌더 확인
package migration

import (
	"strings"
	"testing"
)

func TestCreateTable_SQL_NoPrimaryKey(t *testing.T) {
	op := CreateTable{Table: &Table{
		Name:    "logs",
		Columns: []*Column{{Name: "msg", Type: CanonicalType{Base: "TEXT"}, Nullable: true}},
	}}
	got := op.SQL()
	if strings.Contains(got, "PRIMARY KEY") {
		t.Errorf("SQL() should not contain PRIMARY KEY:\n%s", got)
	}
}

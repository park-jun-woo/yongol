//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestCreateTable_SQL — 컬럼 + PRIMARY KEY 절 렌더 확인
package migration

import (
	"strings"
	"testing"
)

func TestCreateTable_SQL(t *testing.T) {
	op := CreateTable{Table: &Table{
		Name: "users",
		Columns: []*Column{
			{Name: "id", Type: CanonicalType{Base: "BIGINT"}, Nullable: false},
			{Name: "email", Type: CanonicalType{Base: "VARCHAR", Length: 255}, Nullable: true},
		},
		PrimaryKey: []string{"id"},
	}}
	got := op.SQL()
	for _, want := range []string{
		"CREATE TABLE users (",
		"id BIGINT NOT NULL",
		"email VARCHAR(255)",
		"PRIMARY KEY (id)",
		"\n);",
	} {
		if !strings.Contains(got, want) {
			t.Errorf("SQL() missing %q in:\n%s", want, got)
		}
	}
}

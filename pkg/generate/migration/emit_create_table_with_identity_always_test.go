//ff:func feature=migration type=test control=sequence
//ff:what TestEmitCreateTableWithIdentityAlways — CreateTable.SQL 출력에 GENERATED ALWAYS AS IDENTITY 포함
package migration

import (
	"strings"
	"testing"
)

func TestEmitCreateTableWithIdentityAlways(t *testing.T) {
	tbl := &Table{
		Name: "users",
		Columns: []*Column{
			{
				Name:     "id",
				Type:     CanonicalType{Base: "BIGINT"},
				Nullable: false,
				Identity: &IdentitySpec{Always: true},
			},
			{
				Name:     "email",
				Type:     CanonicalType{Base: "TEXT"},
				Nullable: false,
			},
		},
		PrimaryKey: []string{"id"},
	}
	op := CreateTable{Table: tbl}
	sql := op.SQL()
	if !strings.Contains(sql, "id BIGINT NOT NULL GENERATED ALWAYS AS IDENTITY") {
		t.Errorf("SQL should contain canonical IDENTITY clause, got:\n%s", sql)
	}
	if !strings.Contains(sql, "PRIMARY KEY (id)") {
		t.Errorf("SQL should contain table-level PRIMARY KEY, got:\n%s", sql)
	}
}

//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvert* — direct branch coverage for the per-sequence IR converters
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFindTablePK_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{
		{Name: "projects", PrimaryKey: []string{"id"}},
		{Name: "logs"},
	}}
	if got := findTablePK(fs, "PROJECTS"); got != "id" {
		t.Errorf("findTablePK = %q, want id", got)
	}
	if got := findTablePK(fs, "logs"); got != "" {
		t.Errorf("findTablePK no PK = %q, want empty", got)
	}
	if got := findTablePK(fs, "missing"); got != "" {
		t.Errorf("findTablePK missing = %q, want empty", got)
	}
}

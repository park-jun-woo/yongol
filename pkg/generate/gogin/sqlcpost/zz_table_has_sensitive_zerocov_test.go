//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestTableHasSensitiveColumn_ZeroCov — @sensitive 컬럼 유무 분기 직접 호출

package sqlcpost

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestTableHasSensitiveColumn_ZeroCov(t *testing.T) {
	// no sensitive columns.
	plain := ddl.Table{Columns: map[string]ddl.Column{
		"id":   {Name: "id"},
		"name": {Name: "name"},
	}}
	if tableHasSensitiveColumn(plain) {
		t.Errorf("plain table reported sensitive")
	}
	// one sensitive column.
	withSecret := ddl.Table{Columns: map[string]ddl.Column{
		"id": {Name: "id"},
		"pw": {Name: "pw", Sensitive: true},
	}}
	if !tableHasSensitiveColumn(withSecret) {
		t.Errorf("sensitive column not detected")
	}
}

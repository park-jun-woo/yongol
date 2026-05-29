//ff:func feature=gen-gogin type=test control=sequence
//ff:what hasDDL — Fullstack 에 DDL 이 파싱되어 존재하는지 여부

package boot

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestHasDDL(t *testing.T) {
	if hasDDL(nil) {
		t.Errorf("nil fs should be false")
	}
	if hasDDL(&yongol.Fullstack{}) {
		t.Errorf("empty fs should be false")
	}
	withTables := &yongol.Fullstack{DDLTables: []ddl.Table{{Name: "users"}}}
	if !hasDDL(withTables) {
		t.Errorf("fs with DDLTables should be true")
	}
}

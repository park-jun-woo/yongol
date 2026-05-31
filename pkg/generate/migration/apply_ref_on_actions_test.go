//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestApplyRefOnActions(t *testing.T) {
	fk := &ForeignKey{}
	toks := []string{"users(id)", "ON", "DELETE", "SET", "NULL", "ON", "UPDATE", "CASCADE"}
	end := applyRefOnActions(fk, toks, 1)
	if fk.OnDelete != "SET NULL" || fk.OnUpdate != "CASCADE" {
		t.Errorf("OnDelete=%q OnUpdate=%q, want SET NULL / CASCADE", fk.OnDelete, fk.OnUpdate)
	}
	if end != len(toks) {
		t.Errorf("consumed index = %d, want %d", end, len(toks))
	}
}

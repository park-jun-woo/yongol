//ff:func feature=migration type=test control=sequence
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestApplyAlterFKOnActions(t *testing.T) {
	fk := &ForeignKey{}
	applyAlterFKOnActions(fk, " ON DELETE CASCADE ON UPDATE RESTRICT")
	if fk.OnDelete != "CASCADE" || fk.OnUpdate != "RESTRICT" {
		t.Errorf("got OnDelete=%q OnUpdate=%q", fk.OnDelete, fk.OnUpdate)
	}
	fk2 := &ForeignKey{}
	applyAlterFKOnActions(fk2, "")
	if fk2.OnDelete != "" || fk2.OnUpdate != "" {
		t.Errorf("empty tail should leave actions empty: %+v", fk2)
	}
}

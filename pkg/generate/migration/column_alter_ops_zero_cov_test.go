//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestColumnAlterOps_ZeroCov(t *testing.T) {
	prevMap := map[string]*Column{"c": {Name: "c", Type: CanonicalType{Base: "INTEGER"}}}
	currMap := map[string]*Column{"c": {Name: "c", Type: CanonicalType{Base: "BIGINT"}}}
	ops := columnAlterOps("t", []string{"c"}, prevMap, currMap, map[string]bool{}, newEmptyHints())
	if len(ops) == 0 {
		t.Errorf("expected alter ops")
	}
}

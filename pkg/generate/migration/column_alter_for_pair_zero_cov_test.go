//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestColumnAlterForPair_ZeroCov(t *testing.T) {
	pc := &Column{Name: "c", Type: CanonicalType{Base: "INTEGER"}, Nullable: true}
	cc := &Column{Name: "c", Type: CanonicalType{Base: "BIGINT"}, Nullable: false, Default: "0"}
	ops := columnAlterForPair("t", "c", pc, cc, newEmptyHints())
	if len(ops) != 3 {
		t.Errorf("expected type+nullable+default ops, got %d", len(ops))
	}
}

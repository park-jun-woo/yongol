//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
)

func TestDispatchStatement_ZeroCov(t *testing.T) {
	s := NewSchema()
	if err := dispatchStatement(s, "CREATE TABLE t (id BIGINT PRIMARY KEY)"); err != nil {
		t.Fatalf("dispatch: %v", err)
	}
	if _, ok := s.Tables["t"]; !ok {
		t.Errorf("table not created")
	}
	if err := dispatchStatement(s, ""); err != nil {
		t.Errorf("empty stmt should be nil")
	}
}

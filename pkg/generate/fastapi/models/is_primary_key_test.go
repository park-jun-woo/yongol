//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증
package models

import (
	"testing"
)

func TestIsPrimaryKey(t *testing.T) {
	pk := []string{"id", "tenant_id"}
	if !isPrimaryKey("id", pk) {
		t.Errorf("id should be primary key")
	}
	if !isPrimaryKey("tenant_id", pk) {
		t.Errorf("tenant_id should be primary key")
	}
	if isPrimaryKey("name", pk) {
		t.Errorf("name should not be primary key")
	}
	if isPrimaryKey("id", nil) {
		t.Errorf("empty pk should yield false")
	}
}

//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package migration

import (
	"testing"
	"time"
)

func TestEmitMigration_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	curr := NewSchema()
	curr.Tables["t"] = &Table{Name: "t", Columns: []*Column{{Name: "id", Type: CanonicalType{Base: "BIGINT"}}}}
	ops := []Operation{CreateTable{Table: curr.Tables["t"]}}
	res, _, err := emitMigration(dir, dir, curr, ops, newEmptyHints(), ModeInitial, "v0.0.0", time.Unix(0, 0).UTC(), nil)
	if err != nil {
		t.Fatalf("emitMigration: %v", err)
	}
	if res.MigrationFile == "" {
		t.Errorf("no migration file recorded")
	}
}

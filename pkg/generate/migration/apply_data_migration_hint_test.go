//ff:func feature=migration type=test control=sequence
//ff:what apply*Hint 단위 테스트 — 함수명 컨벤션에 맞춘 직접 커버 (전 분기)
package migration

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestApplyDataMigrationHint(t *testing.T) {
	h := newEmptyHints()
	applyDataMigrationHint(h, ddl.HintComment{TableCtx: "users", Args: map[string]string{"file": "u.sql"}})
	if h.DataMigrations["users"] != "u.sql" {
		t.Errorf("data migration not stored: %v", h.DataMigrations)
	}
	// missing file → no-op.
	applyDataMigrationHint(h, ddl.HintComment{TableCtx: "x", Args: map[string]string{}})
	// missing table ctx → no-op.
	applyDataMigrationHint(h, ddl.HintComment{Args: map[string]string{"file": "y.sql"}})
	if len(h.DataMigrations) != 1 {
		t.Errorf("invalid data migration hints should be ignored: %v", h.DataMigrations)
	}
}

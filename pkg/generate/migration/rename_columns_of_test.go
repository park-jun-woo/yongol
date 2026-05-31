//ff:func feature=migration type=test control=sequence
//ff:what 마이그레이션 파이프라인 함수별 named 테스트 — tsma 함수명 매칭용 (parse/diff/emit/tokenizer 커버)
package migration

import (
	"testing"
)

func TestRenameColumnsOf(t *testing.T) {
	prev, _, _ := migPipelineSchemas(t)
	users := prev.Tables["users"]
	rules := []RenameColumnHint{{Table: "users", From: "email", To: "email_addr"}}
	cols := renameColumnsOf(users, "users", rules)
	if len(cols) == 0 {
		t.Errorf("expected columns")
	}
	// no rules → original columns.
	if got := renameColumnsOf(users, "users", nil); len(got) != len(users.Columns) {
		t.Errorf("no rules should return original columns")
	}
}

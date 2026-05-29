//ff:func feature=migration type=test control=sequence
//ff:what TestParseIdentityRejectsDefaultConflict — IDENTITY + DEFAULT 동시 선언 시 파서 에러 반환
package migration

import (
	"strings"
	"testing"
)

func TestParseIdentityRejectsDefaultConflict(t *testing.T) {
	sql := `
CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY DEFAULT 0 PRIMARY KEY,
    email TEXT NOT NULL
);
`
	s := NewSchema()
	err := BuildASTFromSQL(s, sql)
	if err == nil {
		t.Fatalf("expected parse error for IDENTITY + DEFAULT conflict, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "IDENTITY") || !strings.Contains(msg, "DEFAULT") {
		t.Errorf("error message should mention both IDENTITY and DEFAULT: %q", msg)
	}
}

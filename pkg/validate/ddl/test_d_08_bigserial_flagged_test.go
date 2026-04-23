//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD08BigserialFlagged — BIGSERIAL 컬럼에 D-8 ERROR 발생 확인
package ddl

import (
	"strings"
	"testing"
)

func TestD08BigserialFlagged(t *testing.T) {
	sql := `CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    email TEXT NOT NULL
);`
	diags := runD08InTmpDir(t, "users.sql", sql)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	d := diags[0]
	if !strings.Contains(d.Message, "[D-8]") {
		t.Errorf("message should contain [D-8], got %q", d.Message)
	}
	if !strings.Contains(d.Message, "BIGSERIAL") {
		t.Errorf("message should mention BIGSERIAL, got %q", d.Message)
	}
	if !strings.Contains(d.Advice, "BIGINT GENERATED ALWAYS AS IDENTITY") {
		t.Errorf("advice should recommend BIGINT GENERATED ALWAYS AS IDENTITY, got %q", d.Advice)
	}
}

//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD08CaseInsensitive — 소문자 bigserial 도 D-8 에 걸림
package ddl

import (
	"strings"
	"testing"
)

func TestD08CaseInsensitive(t *testing.T) {
	sql := `CREATE TABLE users (
    id bigserial PRIMARY KEY,
    email TEXT NOT NULL
);`
	diags := runD08InTmpDir(t, "users.sql", sql)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic for lowercase bigserial, got %d: %+v", len(diags), diags)
	}
	d := diags[0]
	if !strings.Contains(d.Message, "BIGSERIAL") {
		t.Errorf("message should uppercase the raw type, got %q", d.Message)
	}
}

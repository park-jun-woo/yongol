//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD08IdentityPasses — IDENTITY 컬럼은 D-8 에 걸리지 않음 (false positive 방지)
package ddl

import "testing"

func TestD08IdentityPasses(t *testing.T) {
	sql := `CREATE TABLE users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email TEXT NOT NULL
);`
	diags := runD08InTmpDir(t, "users.sql", sql)
	if len(diags) != 0 {
		t.Fatalf("IDENTITY column must not trigger D-8, got %d diagnostics: %+v", len(diags), diags)
	}
}

//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD08SerialFlagged — SERIAL 컬럼에 D-8 ERROR 발생 + INTEGER 대체 권고
package ddl

import (
	"strings"
	"testing"
)

func TestD08SerialFlagged(t *testing.T) {
	sql := `CREATE TABLE things (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);`
	diags := runD08InTmpDir(t, "things.sql", sql)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	d := diags[0]
	if !strings.Contains(d.Message, "SERIAL") {
		t.Errorf("message should mention SERIAL, got %q", d.Message)
	}
	if !strings.Contains(d.Advice, "INTEGER GENERATED ALWAYS AS IDENTITY") {
		t.Errorf("advice should recommend INTEGER GENERATED ALWAYS AS IDENTITY, got %q", d.Advice)
	}
}

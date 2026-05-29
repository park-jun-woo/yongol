//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD08SmallserialFlagged — SMALLSERIAL 컬럼에 D-8 ERROR + SMALLINT 대체 권고
package ddl

import (
	"strings"
	"testing"
)

func TestD08SmallserialFlagged(t *testing.T) {
	sql := `CREATE TABLE tiny (
    id SMALLSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);`
	diags := runD08InTmpDir(t, "tiny.sql", sql)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	d := diags[0]
	if !strings.Contains(d.Message, "SMALLSERIAL") {
		t.Errorf("message should mention SMALLSERIAL, got %q", d.Message)
	}
	if !strings.Contains(d.Advice, "SMALLINT GENERATED ALWAYS AS IDENTITY") {
		t.Errorf("advice should recommend SMALLINT GENERATED ALWAYS AS IDENTITY, got %q", d.Advice)
	}
}

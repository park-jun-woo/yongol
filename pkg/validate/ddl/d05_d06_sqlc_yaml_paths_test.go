//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"testing"
)

func TestD05D06SqlcYamlPaths(t *testing.T) {
	// good config → no diags
	fs := zcWriteSpecs(t, zcGoodSqlc, "")
	if d := d05SqlcYamlSchemaPath(fs); len(d) != 0 {
		t.Errorf("good schema path want 0, got %d", len(d))
	}
	if d := d06SqlcYamlQueriesPath(fs); len(d) != 0 {
		t.Errorf("good queries path want 0, got %d", len(d))
	}
	// bad paths → warnings
	bad := "version: \"2\"\nsql:\n  - schema: \"models/\"\n    queries: \"sql/\"\n    engine: postgresql\n"
	fsBad := zcWriteSpecs(t, bad, "")
	if d := d05SqlcYamlSchemaPath(fsBad); len(d) != 1 {
		t.Errorf("bad schema path want 1, got %d", len(d))
	}
	if d := d06SqlcYamlQueriesPath(fsBad); len(d) != 1 {
		t.Errorf("bad queries path want 1, got %d", len(d))
	}
}

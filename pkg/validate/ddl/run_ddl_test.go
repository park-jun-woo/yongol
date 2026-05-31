//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"testing"
)

func TestRunDDL(t *testing.T) {
	schema := "CREATE TABLE post (\n" +
		"  id BIGINT PRIMARY KEY,\n" +
		"  author_id BIGINT REFERENCES users(id),\n" +
		"  password_hash TEXT NOT NULL\n" +
		");\n"
	fs := zcWriteSpecs(t, zcGoodSqlc, schema)
	diags := Run(fs)
	// Run aggregates D-15 (nullable FK) and XDD-61 (sensitive) at minimum.
	if len(diags) == 0 {
		t.Errorf("Run want diags, got 0")
	}
}

//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"strings"
	"testing"
)

func TestXdd61SensitiveNoAnnotation_ZeroCov(t *testing.T) {
	schema := "CREATE TABLE account (\n" +
		"  id BIGINT PRIMARY KEY,\n" +
		"  password_hash TEXT NOT NULL\n" +
		");\n"
	fs := zcWriteSpecs(t, zcGoodSqlc, schema)
	d := xdd61SensitiveNoAnnotation(fs)
	if len(d) != 1 || !strings.Contains(d[0].Message, "[XDD-61]") {
		t.Errorf("sensitive col want 1 XDD-61 diag, got %+v", d)
	}
	// annotated → none
	schemaAnn := "CREATE TABLE account (\n" +
		"  id BIGINT PRIMARY KEY,\n" +
		"  password_hash TEXT NOT NULL -- @sensitive\n" +
		");\n"
	fsAnn := zcWriteSpecs(t, zcGoodSqlc, schemaAnn)
	if d := xdd61SensitiveNoAnnotation(fsAnn); len(d) != 0 {
		t.Errorf("annotated want 0, got %d", len(d))
	}
}

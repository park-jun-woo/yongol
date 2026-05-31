//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD15FkNullable(t *testing.T) {
	schema := "CREATE TABLE post (\n" +
		"  id BIGINT PRIMARY KEY,\n" +
		"  author_id BIGINT REFERENCES users(id)\n" +
		");\n"
	fs := zcWriteSpecs(t, zcGoodSqlc, schema)
	d := d15FkNullable(fs)
	if len(d) != 1 || !strings.Contains(d[0].Message, "[D-15]") {
		t.Errorf("nullable FK want 1 D-15 diag, got %+v", d)
	}
	// NOT NULL FK → none
	schemaOK := "CREATE TABLE post (\n" +
		"  id BIGINT PRIMARY KEY,\n" +
		"  author_id BIGINT NOT NULL REFERENCES users(id)\n" +
		");\n"
	fsOK := zcWriteSpecs(t, zcGoodSqlc, schemaOK)
	if d := d15FkNullable(fsOK); len(d) != 0 {
		t.Errorf("NOT NULL FK want 0, got %d", len(d))
	}
	// no db files → none
	if d := d15FkNullable(&yongol.Fullstack{SpecsDir: t.TempDir()}); len(d) != 0 {
		t.Errorf("no files want 0, got %d", len(d))
	}
}

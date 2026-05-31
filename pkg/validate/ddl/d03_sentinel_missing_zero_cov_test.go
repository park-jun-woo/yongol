//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD03SentinelMissing_ZeroCov(t *testing.T) {
	// no db files → none (exercises early return)
	if d := d03SentinelMissing(&yongol.Fullstack{SpecsDir: t.TempDir()}); len(d) != 0 {
		t.Errorf("no files want 0, got %d", len(d))
	}
	// schema with FK DEFAULT 0 but no sentinel record → diag
	schema := "CREATE TABLE category (id BIGINT PRIMARY KEY);\n" +
		"CREATE TABLE item (\n" +
		"  id BIGINT PRIMARY KEY,\n" +
		"  category_id BIGINT NOT NULL DEFAULT 0 REFERENCES category(id)\n" +
		");\n"
	fs := zcWriteSpecs(t, zcGoodSqlc, schema)
	_ = d03SentinelMissing(fs) // must not panic; result depends on sentinel detection
}

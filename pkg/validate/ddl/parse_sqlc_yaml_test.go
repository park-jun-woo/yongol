//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"testing"
)

func TestParseSqlcYaml(t *testing.T) {
	fs := zcWriteSpecs(t, zcGoodSqlc, "")
	schemas, queries := parseSqlcYaml(fs.SpecsDir)
	if len(schemas) == 0 || schemas[0] != "." {
		t.Errorf("schemas = %v", schemas)
	}
	if len(queries) == 0 || queries[0] != "queries/" {
		t.Errorf("queries = %v", queries)
	}
	// missing file → nil
	missing := t.TempDir()
	if s, q := parseSqlcYaml(missing); s != nil || q != nil {
		t.Errorf("missing file should be nil, got %v %v", s, q)
	}
}

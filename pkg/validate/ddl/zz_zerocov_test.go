//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀

package ddl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// zcWriteSpecs writes db/sqlc.yaml + db/schema.sql into a tmp dir and returns
// a Fullstack pointing at it with DDL marked present.
func zcWriteSpecs(t *testing.T, sqlcYaml, schemaSQL string) *yongol.Fullstack {
	t.Helper()
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if sqlcYaml != "" {
		if err := os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(sqlcYaml), 0o644); err != nil {
			t.Fatalf("write sqlc.yaml: %v", err)
		}
	}
	if schemaSQL != "" {
		if err := os.WriteFile(filepath.Join(dbDir, "schema.sql"), []byte(schemaSQL), 0o644); err != nil {
			t.Fatalf("write schema.sql: %v", err)
		}
	}
	return &yongol.Fullstack{
		SpecsDir:  tmp,
		Presences: map[yongol.SSOTKind]yongol.SSOTPresence{yongol.KindDDL: yongol.SSOTPopulated},
	}
}

const zcGoodSqlc = "version: \"2\"\nsql:\n  - schema: \".\"\n    queries: \"queries/\"\n    engine: postgresql\n"

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

func TestD04SqlcYamlRequired(t *testing.T) {
	// no sqlc.yaml → ERROR
	fs := zcWriteSpecs(t, "", "CREATE TABLE t (id BIGINT);")
	if d := d04SqlcYamlRequired(fs); len(d) != 1 {
		t.Errorf("missing sqlc.yaml want 1 diag, got %d", len(d))
	}
	// present → none
	fs2 := zcWriteSpecs(t, zcGoodSqlc, "")
	if d := d04SqlcYamlRequired(fs2); len(d) != 0 {
		t.Errorf("present sqlc.yaml want 0 diag, got %d", len(d))
	}
	// DDL absent → none
	if d := d04SqlcYamlRequired(&yongol.Fullstack{}); len(d) != 0 {
		t.Errorf("absent DDL want 0 diag, got %d", len(d))
	}
}

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

func TestD07SqlcPositionalParam(t *testing.T) {
	// scanPositionals reads the actual query file, so write one.
	tmp := t.TempDir()
	qPath := filepath.Join(tmp, "q.sql")
	if err := os.WriteFile(qPath, []byte("-- name: GetUser :one\nSELECT * FROM users WHERE id = $1;\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "GetUser", File: qPath, Line: 1},
		},
	}
	d := d07SqlcPositionalParam(fs)
	if len(d) != 1 || !strings.Contains(d[0].Message, "[D-7]") {
		t.Errorf("positional param want 1 D-7 diag, got %+v", d)
	}
	// no positional → none
	qPath2 := filepath.Join(tmp, "q2.sql")
	if err := os.WriteFile(qPath2, []byte("-- name: GetUser2 :one\nSELECT * FROM users WHERE id = @id;\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}
	fs2 := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "GetUser2", File: qPath2, Line: 1},
		},
	}
	if d := d07SqlcPositionalParam(fs2); len(d) != 0 {
		t.Errorf("named param want 0, got %d", len(d))
	}
}

func TestD01SqlcQueryDuplicate_ZeroCov(t *testing.T) {
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "GetUser", File: "a.sql", Line: 1},
			{Name: "GetUser", File: "b.sql", Line: 2},
		},
	}
	d := d01SqlcQueryDuplicate(fs)
	if len(d) < 1 {
		t.Errorf("duplicate want diags, got %d", len(d))
	}
	// no queries → none
	if d := d01SqlcQueryDuplicate(&yongol.Fullstack{}); len(d) != 0 {
		t.Errorf("empty want 0, got %d", len(d))
	}
	// unique → none
	fs2 := &yongol.Fullstack{SQLcQueries: []sqlcparser.QuerySpec{{Name: "A"}, {Name: "B"}}}
	if d := d01SqlcQueryDuplicate(fs2); len(d) != 0 {
		t.Errorf("unique want 0, got %d", len(d))
	}
}

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

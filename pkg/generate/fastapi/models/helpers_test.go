//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestModelsHelpers — fastapi models 패키지 순수 헬퍼(타입 매핑·PK·FK·기본값·table_args) 검증

package models

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestIsPrimaryKey(t *testing.T) {
	pk := []string{"id", "tenant_id"}
	if !isPrimaryKey("id", pk) {
		t.Errorf("id should be primary key")
	}
	if !isPrimaryKey("tenant_id", pk) {
		t.Errorf("tenant_id should be primary key")
	}
	if isPrimaryKey("name", pk) {
		t.Errorf("name should not be primary key")
	}
	if isPrimaryKey("id", nil) {
		t.Errorf("empty pk should yield false")
	}
}

func TestPyFamily(t *testing.T) {
	cases := map[string]string{
		"BIGINT":           "int",
		"INT2":             "int",
		"TEXT":             "str",
		"CITEXT":           "str",
		"BOOLEAN":          "bool",
		"TIMESTAMPTZ":      "datetime",
		"DATE":             "date",
		"UUID":             "uuid.UUID",
		"JSONB":            "dict[str, Any]",
		"NUMERIC":          "Decimal",
		"DOUBLE PRECISION": "float",
		"BYTEA":            "bytes",
		"INET":             "str",
		"INTERVAL":         "timedelta",
		"UNKNOWNTYPE":      "str",
	}
	for in, want := range cases {
		if got := pyFamily(in); got != want {
			t.Errorf("pyFamily(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestMapSAFamily(t *testing.T) {
	cases := map[string]string{
		"BIGINT":                      "Integer",
		"INTEGER":                     "Integer",
		"TEXT":                        "Text",
		"VARCHAR":                     "String",
		"BOOLEAN":                     "Boolean",
		"TIMESTAMP":                   "DateTime(timezone=False)",
		"TIMESTAMP WITHOUT TIME ZONE": "DateTime(timezone=False)",
		"TIMESTAMPTZ":                 "DateTime(timezone=True)",
		"DATE":                        "Date",
		"UUID":                        "Uuid",
		"JSONB":                       "JSONB",
		"NUMERIC":                     "Numeric",
		"REAL":                        "Float",
		"BYTEA":                       "LargeBinary",
		"CIDR":                        "INET",
		"INTERVAL":                    "Interval",
		"WEIRD":                       "String",
	}
	for in, want := range cases {
		if got := mapSAFamily(in); got != want {
			t.Errorf("mapSAFamily(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestPascalCase(t *testing.T) {
	cases := map[string]string{
		"users":       "Users",
		"order_items": "OrderItems",
		"":            "",
		"_leading":    "Leading",
		"a__b":        "AB",
		"single":      "Single",
		"trailing_":   "Trailing",
	}
	for in, want := range cases {
		if got := pascalCase(in); got != want {
			t.Errorf("pascalCase(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestMapPGToPython(t *testing.T) {
	cases := []struct {
		raw     string
		notNull bool
		want    string
	}{
		{"BIGINT", true, "int"},
		{"BIGINT", false, "int | None"},
		{"VARCHAR(255)", true, "str"},
		{"TEXT[]", true, "list[str]"},
		{"INTEGER[]", false, "list[int] | None"},
		{"timestamptz", true, "datetime"},
	}
	for _, c := range cases {
		if got := mapPGToPython(c.raw, c.notNull); got != c.want {
			t.Errorf("mapPGToPython(%q,%v) = %q, want %q", c.raw, c.notNull, got, c.want)
		}
	}
}

func TestMapPGToSA(t *testing.T) {
	cases := map[string]string{
		"BIGINT":       "Integer",
		"VARCHAR(255)": "String",
		"TEXT[]":       "ARRAY(Text)",
		"uuid":         "Uuid",
	}
	for in, want := range cases {
		if got := mapPGToSA(in); got != want {
			t.Errorf("mapPGToSA(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestFindForeignKeyAttr(t *testing.T) {
	fks := []ddl.ForeignKey{
		{Column: "user_id", RefTable: "users", RefColumn: "id"},
	}
	if got := findForeignKeyAttr(fks, "user_id"); got != `ForeignKey("users.id")` {
		t.Errorf("unexpected fk attr: %q", got)
	}
	if got := findForeignKeyAttr(fks, "other"); got != "" {
		t.Errorf("expected empty for non-matching col, got %q", got)
	}
	if got := findForeignKeyAttr(nil, "user_id"); got != "" {
		t.Errorf("expected empty for nil fks, got %q", got)
	}
}

func TestFormatColumnList(t *testing.T) {
	if got := formatColumnList([]string{"a", "b"}); got != `"a", "b"` {
		t.Errorf("unexpected: %q", got)
	}
	if got := formatColumnList([]string{"only"}); got != `"only"` {
		t.Errorf("unexpected single: %q", got)
	}
	if got := formatColumnList(nil); got != "" {
		t.Errorf("expected empty for nil, got %q", got)
	}
}

func TestSaDefault(t *testing.T) {
	cases := []struct {
		col  ddl.Column
		want string
	}{
		{ddl.Column{RawType: "UUID", HasDefault: true}, "default=uuid.uuid4"},
		{ddl.Column{RawType: "TIMESTAMPTZ", HasDefault: true}, `server_default="now()"`},
		{ddl.Column{RawType: "DATE", HasDefault: true}, `server_default="now()"`},
		{ddl.Column{RawType: "SERIAL"}, ""},
		{ddl.Column{RawType: "BIGSERIAL"}, ""},
		{ddl.Column{RawType: "TEXT", DefaultLiteral: ""}, ""},
		{ddl.Column{RawType: "BOOLEAN", HasDefault: true, DefaultLiteral: "true"}, "default=True"},
		{ddl.Column{RawType: "BOOLEAN", HasDefault: true, DefaultLiteral: "FALSE"}, "default=False"},
		{ddl.Column{RawType: "INTEGER", HasDefault: true, DefaultLiteral: "42"}, "default=42"},
		{ddl.Column{RawType: "TEXT", HasDefault: true, DefaultLiteral: "hello"}, `default="hello"`},
	}
	for _, c := range cases {
		if got := saDefault(c.col); got != c.want {
			t.Errorf("saDefault(%+v) = %q, want %q", c.col, got, c.want)
		}
	}
}

func TestRenderTableArgs(t *testing.T) {
	// No indexes -> empty.
	if got := renderTableArgs(ddl.Table{}); got != "" {
		t.Errorf("expected empty for no indexes, got %q", got)
	}
	// Unique + non-unique index.
	tbl := ddl.Table{
		Indexes: []ddl.Index{
			{Name: "uq_email", Columns: []string{"email"}, IsUnique: true},
			{Name: "idx_name", Columns: []string{"first", "last"}, IsUnique: false},
		},
	}
	got := renderTableArgs(tbl)
	if !strings.Contains(got, `UniqueConstraint("email", name="uq_email")`) {
		t.Errorf("missing unique constraint: %q", got)
	}
	if !strings.Contains(got, `Index("idx_name", "first", "last")`) {
		t.Errorf("missing index: %q", got)
	}
	if !strings.HasPrefix(got, "    __table_args__ = (\n") || !strings.HasSuffix(got, "    )\n") {
		t.Errorf("malformed table_args wrapper: %q", got)
	}
}

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestZeroCov — 0% util/adapter 함수 (bindGoType / GoTypeOf / registry / family dispatch / pgtype) 회귀

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
	"github.com/park-jun-woo/yongol/pkg/generate/typemap"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestBindingKindString(t *testing.T) {
	cases := []struct {
		k    BindingKind
		want string
	}{
		{KindNative, "Native"},
		{KindPointer, "Pointer"},
		{KindPgtype, "Pgtype"},
		{KindJSONB, "JSONB"},
		{KindBytea, "Bytea"},
		{KindArray, "Array"},
		{KindEnum, "Enum"},
		{KindUnsupported, "Unsupported"},
		{BindingKind(99), "Unknown"},
	}
	for _, c := range cases {
		if got := c.k.String(); got != c.want {
			t.Errorf("BindingKind(%d).String() = %q, want %q", c.k, got, c.want)
		}
	}
}

func TestBindGoType_AllFamilies(t *testing.T) {
	families := []typemap.PGFamily{
		typemap.FamilyEnum,
		typemap.FamilyArray,
		typemap.FamilyUUID,
		typemap.FamilyNumeric,
		typemap.FamilyTimestampTZ,
		typemap.FamilyTimestamp,
		typemap.FamilyDate,
		typemap.FamilyInet,
		typemap.FamilyInterval,
		typemap.FamilyJSONB,
		typemap.FamilyBytea,
		typemap.FamilyInteger,
		typemap.FamilyFloat,
		typemap.FamilyString,
		typemap.FamilyBoolean,
		typemap.FamilyUnsupported,
	}
	for _, f := range families {
		opts := ir.BindOpts{NotNull: true, ElementHead: "TEXT"}
		gb := bindGoType(f, opts)
		_ = gb // every branch must return a value
	}
	// nullable variant to hit non-notNull paths
	for _, f := range families {
		gb := bindGoType(f, ir.BindOpts{NotNull: false, ElementHead: "TEXT"})
		_ = gb
	}
	// unsupported must be flagged not supported
	if bindGoType(typemap.FamilyUnsupported, ir.BindOpts{}).Supported {
		t.Errorf("FamilyUnsupported should not be Supported")
	}
}

func TestGoTypeOf_Matrix(t *testing.T) {
	cases := []struct {
		raw  string
		want string
	}{
		{"BIGINT", "int64"},
		{"INTEGER", "int64"},
		{"SMALLINT", "int64"},
		{"SERIAL", "int64"},
		{"INT4", "int64"},
		{"VARCHAR(50)", "string"},
		{"TEXT", "string"},
		{"UUID", "string"},
		{"CHAR(2)", "string"},
		{"BOOLEAN", "bool"},
		{"BOOL", "bool"},
		{"TIMESTAMPTZ", "time.Time"},
		{"TIMESTAMP", "time.Time"},
		{"DATE", "time.Time"},
		{"NUMERIC(10,2)", "float64"},
		{"REAL", "float64"},
		{"FLOAT8", "float64"},
		{"JSONB", "json.RawMessage"},
		{"JSON", "json.RawMessage"},
		{"TEXT[]", "string"},     // array suffix stripped → TEXT
		{"BIGINT[]", "int64"},    // array suffix stripped → BIGINT
		{"MYSTERYTYPE", "string"}, // fallthrough default
	}
	for _, c := range cases {
		got := GoTypeOf(ddl.Column{RawType: c.raw})
		if got != c.want {
			t.Errorf("GoTypeOf(%q) = %q, want %q", c.raw, got, c.want)
		}
	}
}

func TestGoTypeRegistry_BindAndFactory(t *testing.T) {
	r := NewGoTypeRegistry()
	if r == nil {
		t.Fatal("NewGoTypeRegistry returned nil")
	}
	tb := r.Bind(typemap.FamilyInteger, ir.BindOpts{NotNull: true})
	if !tb.Supported {
		t.Errorf("integer NOT NULL should be supported")
	}
	if tb.Family != typemap.FamilyInteger {
		t.Errorf("Family = %v, want FamilyInteger", tb.Family)
	}
	if tb.DBType != "int64" {
		t.Errorf("DBType = %q, want int64", tb.DBType)
	}
}

func TestToIRBinding_Mapping(t *testing.T) {
	gb := GoTypeBinding{
		SqlcGoType:     "pgtype.Int4",
		ApiField:       "*int64",
		ConvertExpr:    "conv",
		InsertExpr:     "ins",
		ResponseExpr:   "resp",
		NilCheckExpr:   "nil",
		Imports:        []string{"pkg/a"},
		Supported:      true,
		DefaultLiteral: "0",
	}
	opts := ir.BindOpts{NotNull: false}
	got := toIRBinding(typemap.FamilyInteger, opts, gb)
	if got.DBType != "pgtype.Int4" || got.APIType != "*int64" {
		t.Errorf("type mapping mismatch: %+v", got)
	}
	if got.ToDBExpr != "ins" || got.ToAPIExpr != "conv" || got.ToResponseExpr != "resp" {
		t.Errorf("expr mapping mismatch: %+v", got)
	}
	if got.NotNull != false || got.Supported != true || got.DefaultLiteral != "0" {
		t.Errorf("scalar mapping mismatch: %+v", got)
	}
	if len(got.APIImports) != 1 || got.APIImports[0] != "pkg/a" {
		t.Errorf("imports mapping mismatch: %+v", got.APIImports)
	}
}

func TestMapNativeFamily_Branches(t *testing.T) {
	if _, ok := mapNativeFamily("BIGINT", true, ""); !ok {
		t.Errorf("BIGINT should be native")
	}
	if _, ok := mapNativeFamily("REAL", true, ""); !ok {
		t.Errorf("REAL should be native")
	}
	if _, ok := mapNativeFamily("VARCHAR", true, ""); !ok {
		t.Errorf("VARCHAR should be native")
	}
	if _, ok := mapNativeFamily("BOOLEAN", true, ""); !ok {
		t.Errorf("BOOLEAN should be native")
	}
	if _, ok := mapNativeFamily("UUID", true, ""); ok {
		t.Errorf("UUID should not be native")
	}
}

func TestMapPgtypeFamily_Branches(t *testing.T) {
	heads := []string{"UUID", "NUMERIC", "DECIMAL", "TIMESTAMPTZ", "TIMESTAMP", "DATE", "INET", "CIDR", "INTERVAL", "JSONB", "JSON", "BYTEA"}
	for _, h := range heads {
		if _, ok := mapPgtypeFamily(h, true, ""); !ok {
			t.Errorf("%q should be pgtype family", h)
		}
	}
	if _, ok := mapPgtypeFamily("BIGINT", true, ""); ok {
		t.Errorf("BIGINT should not be pgtype family")
	}
}

func TestNativeFloat_Branches(t *testing.T) {
	nn := nativeFloat(true, "0")
	if nn.SqlcGoType != "float64" || nn.Kind != KindNative {
		t.Errorf("NOT NULL float = %+v", nn)
	}
	nullable := nativeFloat(false, "0")
	if nullable.Kind != KindPgtype {
		t.Errorf("nullable float should be pgtype, got %+v", nullable)
	}
}

func TestPgtypeInt2Int4(t *testing.T) {
	i2 := pgtypeInt2("0")
	if i2.SqlcGoType != "pgtype.Int2" || i2.Kind != KindPgtype {
		t.Errorf("pgtypeInt2 = %+v", i2)
	}
	i4 := pgtypeInt4("0")
	if i4.SqlcGoType != "pgtype.Int4" || i4.Kind != KindPgtype {
		t.Errorf("pgtypeInt4 = %+v", i4)
	}
}

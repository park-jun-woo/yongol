//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestHelpersZeroCov — head 판정 / native·pgtype·array·enum·jsonb·bytea binding 헬퍼 직접 커버

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestHeadPredicates_ZeroCov(t *testing.T) {
	if !isBooleanHead("BOOLEAN") || isBooleanHead("TEXT") {
		t.Errorf("isBooleanHead mismatch")
	}
	if !isFloatHead("REAL") || isFloatHead("TEXT") {
		t.Errorf("isFloatHead mismatch")
	}
	if !isIntegerHead("BIGINT") || isIntegerHead("TEXT") {
		t.Errorf("isIntegerHead mismatch")
	}
	if !isStringHead("TEXT") || isStringHead("BIGINT") {
		t.Errorf("isStringHead mismatch")
	}
}

func TestIsEffectivelyNotNull_ZeroCov(t *testing.T) {
	if isEffectivelyNotNull(ddl.Column{NullableAnnot: true, NotNull: true}) {
		t.Errorf("@nullable annotation must override NOT NULL")
	}
	if !isEffectivelyNotNull(ddl.Column{NotNull: true}) {
		t.Errorf("NOT NULL without annotation must be true")
	}
	if isEffectivelyNotNull(ddl.Column{NotNull: false}) {
		t.Errorf("nullable column must be false")
	}
}

func TestColumnAdapter_ZeroCov(t *testing.T) {
	a := columnAdapter{ddl.Column{RawType: "VARCHAR(50)", CheckEnum: []string{"a", "b"}}}
	if a.RawType() != "VARCHAR(50)" {
		t.Errorf("RawType() = %q", a.RawType())
	}
	if got := a.CheckEnum(); len(got) != 2 || got[0] != "a" {
		t.Errorf("CheckEnum() = %v", got)
	}
}

func TestNativeBindings_ZeroCov(t *testing.T) {
	// NOT NULL native forms.
	if b := nativeInteger(true, "0"); b.SqlcGoType != "int64" || b.Kind != KindNative {
		t.Errorf("nativeInteger NOT NULL = %+v", b)
	}
	// Nullable → pgtype.
	if b := nativeInteger(false, "0"); b.Kind != KindPgtype {
		t.Errorf("nativeInteger nullable = %+v", b)
	}
	if b := nativeString(true, "''"); b.SqlcGoType != "string" || b.Kind != KindNative {
		t.Errorf("nativeString NOT NULL = %+v", b)
	}
	if b := nativeString(false, "''"); b.Kind != KindPgtype {
		t.Errorf("nativeString nullable = %+v", b)
	}
	if b := nativeBoolean(true, "false"); b.SqlcGoType != "bool" || b.Kind != KindNative {
		t.Errorf("nativeBoolean NOT NULL = %+v", b)
	}
	if b := nativeBoolean(false, "false"); b.Kind != KindPgtype {
		t.Errorf("nativeBoolean nullable = %+v", b)
	}
}

func TestNativeFloatWithHead_ZeroCov(t *testing.T) {
	if b := nativeFloatWithHead("REAL", true, "0"); b.SqlcGoType != "float64" || b.Kind != KindNative {
		t.Errorf("float NOT NULL = %+v", b)
	}
	if b := nativeFloatWithHead("REAL", false, "0"); b.SqlcGoType != "pgtype.Float4" {
		t.Errorf("REAL nullable should be Float4: %+v", b)
	}
	if b := nativeFloatWithHead("FLOAT4", false, "0"); b.SqlcGoType != "pgtype.Float4" {
		t.Errorf("FLOAT4 nullable should be Float4: %+v", b)
	}
	if b := nativeFloatWithHead("FLOAT8", false, "0"); b.SqlcGoType != "pgtype.Float8" {
		t.Errorf("FLOAT8 nullable should be Float8: %+v", b)
	}
}

func TestEnumJsonbBytea_ZeroCov(t *testing.T) {
	if b := enumBinding(true, "'pending'"); b.Kind != KindEnum || b.ApiField != "string" {
		t.Errorf("enum NOT NULL = %+v", b)
	}
	if b := enumBinding(false, "'pending'"); b.ApiField != "*string" {
		t.Errorf("enum nullable = %+v", b)
	}
	if b := jsonbBinding(true, "'{}'"); b.Kind != KindJSONB || b.ApiField != "map[string]interface{}" {
		t.Errorf("jsonb NOT NULL = %+v", b)
	}
	if b := jsonbBinding(false, "'{}'"); b.ApiField != "*map[string]interface{}" {
		t.Errorf("jsonb nullable = %+v", b)
	}
	if b := byteaBinding(true, "''"); b.Kind != KindBytea || b.SqlcGoType != "[]byte" {
		t.Errorf("bytea = %+v", b)
	}
	if b := byteaBinding(false, "''"); b.Kind != KindBytea {
		t.Errorf("bytea nullable = %+v", b)
	}
}

func TestArrayBindings_ZeroCov(t *testing.T) {
	// arrayElementGoType across all four supported families + unsupported.
	cases := []struct {
		head string
		want string
		ok   bool
	}{
		{"BIGINT", "int64", true},
		{"REAL", "float64", true},
		{"TEXT", "string", true},
		{"BOOLEAN", "bool", true},
		{"UUID", "", false},
	}
	for _, c := range cases {
		got, ok := arrayElementGoType(c.head)
		if got != c.want || ok != c.ok {
			t.Errorf("arrayElementGoType(%q) = (%q,%v), want (%q,%v)", c.head, got, ok, c.want, c.ok)
		}
	}
	// arrayBinding supported → KindArray.
	if b := arrayBinding("TEXT", "'{}'"); b.Kind != KindArray || b.SqlcGoType != "[]string" {
		t.Errorf("arrayBinding TEXT = %+v", b)
	}
	// arrayBinding unsupported element → KindUnsupported.
	if b := arrayBinding("UUID", "'{}'"); b.Kind != KindUnsupported || b.Supported {
		t.Errorf("arrayBinding UUID should be unsupported: %+v", b)
	}
	// composeArrayBinding directly: supported false branch.
	if b := composeArrayBinding("", false, "UUID", "'{}'"); b.Supported {
		t.Errorf("composeArrayBinding unsupported should not be Supported")
	}
	if b := composeArrayBinding("int64", true, "BIGINT", "'{}'"); b.SqlcGoType != "[]int64" {
		t.Errorf("composeArrayBinding BIGINT = %+v", b)
	}
}

func TestDispatchBinding_ZeroCov(t *testing.T) {
	// Drive dispatchBinding across each family via MapPGType-equivalent columns.
	cases := []struct {
		raw      string
		check    []string
		wantKind BindingKind
	}{
		{"VARCHAR(20)", []string{"a", "b"}, KindEnum},
		{"TEXT[]", nil, KindArray},
		{"UUID", nil, KindPgtype},
		{"NUMERIC(10,2)", nil, KindPgtype},
		{"TIMESTAMPTZ", nil, KindPgtype},
		{"TIMESTAMP", nil, KindPgtype},
		{"DATE", nil, KindPgtype},
		{"INET", nil, KindPgtype},
		{"INTERVAL", nil, KindPgtype},
		{"JSONB", nil, KindJSONB},
		{"BYTEA", nil, KindBytea},
		{"BIGINT", nil, KindNative},
		{"REAL", nil, KindNative},
		{"TEXT", nil, KindNative},
		{"BOOLEAN", nil, KindNative},
	}
	for _, c := range cases {
		col := ddl.Column{RawType: c.raw, NotNull: true, CheckEnum: c.check}
		info := parseRawType(c.raw)
		b := dispatchBinding(col, info, true, "")
		if b.Kind != c.wantKind {
			t.Errorf("dispatchBinding(%q) Kind = %v, want %v", c.raw, b.Kind, c.wantKind)
		}
	}
}

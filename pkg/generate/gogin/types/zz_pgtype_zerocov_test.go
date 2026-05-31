//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestPgtypeConstructorsZeroCov — 모든 pgtype 생성자 + unsupportedBinding 직접 커버

package types

import "testing"

func TestPgtypeScalarConstructors_ZeroCov(t *testing.T) {
	cases := []struct {
		name string
		got  GoTypeBinding
		sqlc string
	}{
		{"bool", pgtypeBool("false"), "pgtype.Bool"},
		{"float4", pgtypeFloat4("0"), "pgtype.Float4"},
		{"float8", pgtypeFloat8("0"), "pgtype.Float8"},
		{"int8", pgtypeInt8("0"), "pgtype.Int8"},
		{"text", pgtypeText("''"), "pgtype.Text"},
	}
	for _, c := range cases {
		if c.got.SqlcGoType != c.sqlc {
			t.Errorf("%s SqlcGoType = %q, want %q", c.name, c.got.SqlcGoType, c.sqlc)
		}
		if c.got.Kind != KindPgtype || !c.got.Supported {
			t.Errorf("%s should be supported pgtype: %+v", c.name, c.got)
		}
	}
}

func TestPgtypeNullableConstructors_ZeroCov(t *testing.T) {
	// inet / interval / numeric / uuid each branch on notNull for the api field.
	if b := pgtypeInet(true, ""); b.ApiField != "string" {
		t.Errorf("pgtypeInet NOT NULL api = %q", b.ApiField)
	}
	if b := pgtypeInet(false, ""); b.ApiField != "*string" {
		t.Errorf("pgtypeInet nullable api = %q", b.ApiField)
	}
	if b := pgtypeInterval(true, ""); b.ApiField != "string" {
		t.Errorf("pgtypeInterval NOT NULL api = %q", b.ApiField)
	}
	if b := pgtypeInterval(false, ""); b.ApiField != "*string" {
		t.Errorf("pgtypeInterval nullable api = %q", b.ApiField)
	}
	if b := pgtypeNumeric(true, ""); b.ApiField != "string" {
		t.Errorf("pgtypeNumeric NOT NULL api = %q", b.ApiField)
	}
	if b := pgtypeNumeric(false, ""); b.ApiField != "*string" {
		t.Errorf("pgtypeNumeric nullable api = %q", b.ApiField)
	}
	if b := pgtypeUUID(true, ""); b.ApiField != "openapi_types.UUID" {
		t.Errorf("pgtypeUUID NOT NULL api = %q", b.ApiField)
	}
	if b := pgtypeUUID(false, ""); b.ApiField != "*openapi_types.UUID" {
		t.Errorf("pgtypeUUID nullable api = %q", b.ApiField)
	}
}

func TestPgtypeTimestamp_ZeroCov(t *testing.T) {
	cases := []struct {
		head string
		sqlc string
	}{
		{"TIMESTAMPTZ", "pgtype.Timestamptz"},
		{"TIMESTAMP", "pgtype.Timestamp"},
		{"DATE", "pgtype.Date"},
	}
	for _, c := range cases {
		if b := pgtypeTimestamp(c.head, true, ""); b.SqlcGoType != c.sqlc {
			t.Errorf("pgtypeTimestamp(%q) NOT NULL = %q, want %q", c.head, b.SqlcGoType, c.sqlc)
		}
		// nullable path appends Ptr to the convert funcs and uses *time.Time.
		if b := pgtypeTimestamp(c.head, false, ""); b.ApiField != "*time.Time" {
			t.Errorf("pgtypeTimestamp(%q) nullable api = %q", c.head, b.ApiField)
		}
	}
}

func TestUnsupportedBinding_ZeroCov(t *testing.T) {
	b := unsupportedBinding("nope")
	if b.Supported || b.Kind != KindUnsupported {
		t.Errorf("unsupportedBinding should be unsupported: %+v", b)
	}
}

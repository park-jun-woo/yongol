//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestHelpersZeroCov — head 판정 / native·pgtype·array·enum·jsonb·bytea binding 헬퍼 직접 커버
package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

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

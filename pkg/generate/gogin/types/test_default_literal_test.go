//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMapPGType_DefaultLiteralPropagated — HasDefault 컬럼의 DefaultLiteral 가 binding 으로 전달되는지

package types

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

func TestMapPGType_DefaultLiteralPropagated(t *testing.T) {
	col := ddl.Column{
		RawType:        "VARCHAR(20)",
		NotNull:        true,
		HasDefault:     true,
		DefaultLiteral: "'member'",
		VarcharLen:     20,
		CheckEnum:      []string{"member", "admin"},
	}
	b := MapPGType(col)
	if b.DefaultLiteral != "'member'" {
		t.Errorf("DefaultLiteral = %q, want %q", b.DefaultLiteral, "'member'")
	}
	if b.Kind != KindEnum {
		t.Errorf("Kind = %v, want KindEnum", b.Kind)
	}
}

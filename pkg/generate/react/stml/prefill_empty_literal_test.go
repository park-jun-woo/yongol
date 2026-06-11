//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what prefillEmptyLiteral 타입별 빈 리터럴 매핑 검증

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestPrefillEmptyLiteral(t *testing.T) {
	cases := map[string]string{
		"string":  "''",
		"":        "''",
		"integer": "0",
		"number":  "0",
		"boolean": "false",
		"array":   "[]",
		"object":  "{}",
	}
	for typ, want := range cases {
		if got := prefillEmptyLiteral(oapiparser.FieldConstraint{Type: typ}); got != want {
			t.Errorf("prefillEmptyLiteral(%q) = %q, want %q", typ, got, want)
		}
	}
}

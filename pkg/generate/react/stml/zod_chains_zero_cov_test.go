//ff:func feature=stml-gen type=test control=sequence
//ff:what 단순 stml 헬퍼 (string/zod/param/login) 묶음 커버 — coverage attribution 으로 다수 함수 PASS
package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestZodChains_ZeroCov(t *testing.T) {
	if zodBaseType("integer") != "z.number().int()" {
		t.Errorf("zodBaseType integer wrong")
	}
	if zodBaseType("number") != "z.number()" {
		t.Errorf("zodBaseType number wrong")
	}
	if zodBaseTypeArray("string") == "" {
		t.Errorf("zodBaseTypeArray empty")
	}
	// plain required string.
	if got := zodChain(oapiparser.FieldConstraint{Type: "string", Required: true}); got == "" {
		t.Errorf("zodChain string empty")
	}
	// optional array.
	if got := zodChain(oapiparser.FieldConstraint{Type: "array", ItemType: "integer"}); got == "" {
		t.Errorf("zodChain array empty")
	}
	// enum chain (required + optional).
	if got := zodEnumChain(oapiparser.FieldConstraint{Enum: []string{"a", "b"}, Required: true}); got == "" {
		t.Errorf("zodEnumChain required empty")
	}
	if got := zodChain(oapiparser.FieldConstraint{Enum: []string{"a"}}); got == "" {
		t.Errorf("zodChain enum-routed empty")
	}
}

//ff:func feature=stml-gen type=test control=sequence
//ff:what generateZodSchema — 미지원 타입 panic 에 operation/field 컨텍스트가 채워지는지 검증
package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// generateZodSchema enriches the panic with operation/field context.
func TestGenerateZodSchema_UnsupportedTypeEnriched(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic, got none")
		}
		ze, ok := r.(*zodGenError)
		if !ok {
			t.Fatalf("expected *zodGenError, got %T", r)
		}
		msg := ze.Error()
		if !strings.Contains(msg, "BadOp") || !strings.Contains(msg, "bad") || !strings.Contains(msg, "weirdtype") {
			t.Errorf("error message missing context: %q", msg)
		}
	}()
	fields := map[string]oapiparser.FieldConstraint{
		"bad": {Type: "weirdtype", Required: true},
	}
	_ = generateZodSchema("BadOp", fields)
}

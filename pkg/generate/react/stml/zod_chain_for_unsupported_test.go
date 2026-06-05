//ff:func feature=stml-gen type=test control=sequence
//ff:what zodChainFor — 미지원 타입 panic 시 *zodGenError 에 operation/field 컨텍스트 채움 검증

package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// An unsupported type panics with *zodGenError enriched with op id + field name.
func TestZodChainFor_UnsupportedEnriched(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic on unsupported type, got none")
		}
		ze, ok := r.(*zodGenError)
		if !ok {
			t.Fatalf("expected *zodGenError, got %T", r)
		}
		if ze.OperationID != "BadOp" {
			t.Errorf("OperationID = %q, want \"BadOp\"", ze.OperationID)
		}
		if ze.Field != "bad" {
			t.Errorf("Field = %q, want \"bad\"", ze.Field)
		}
		if ze.Type != "weirdtype" {
			t.Errorf("Type = %q, want \"weirdtype\"", ze.Type)
		}
		msg := ze.Error()
		if !strings.Contains(msg, "BadOp") || !strings.Contains(msg, "bad") {
			t.Errorf("error message missing context: %q", msg)
		}
	}()
	_ = zodChainFor("BadOp", "bad", oapiparser.FieldConstraint{Type: "weirdtype", Required: true})
}

//ff:func feature=stml-gen type=test control=sequence
//ff:what TestBindCtxField — bindCtx.field 의 정상 조회와 nil 맵/미상 op·필드의 zero-value 반환 검증

package stml

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestBindCtxField(t *testing.T) {
	ctx := bindCtx{
		all: map[string]map[string]oapiparser.FieldTypeInfo{
			"Op": {"flag": {Type: "boolean"}},
		},
		opID: "Op",
	}
	if got := ctx.field("flag"); got.Type != "boolean" {
		t.Errorf("field(flag): want boolean, got %+v", got)
	}
	// Unknown field within the op → zero value.
	if got := ctx.field("missing"); got != (oapiparser.FieldTypeInfo{}) {
		t.Errorf("field(missing): want zero, got %+v", got)
	}
	// Unknown op scope → zero value.
	ctx.opID = "Other"
	if got := ctx.field("flag"); got != (oapiparser.FieldTypeInfo{}) {
		t.Errorf("field(flag) wrong op: want zero, got %+v", got)
	}
	// Nil map → zero value, no panic.
	if got := (bindCtx{}).field("x"); got != (oapiparser.FieldTypeInfo{}) {
		t.Errorf("zero bindCtx field: want zero, got %+v", got)
	}
}

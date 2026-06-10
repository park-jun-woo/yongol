//ff:func feature=stml-gen type=test control=sequence
//ff:what itemParamIsNumber — item 필드 타입(integer/number/string/미상)·비-item 소스·점 경로 분기 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestItemParamIsNumber(t *testing.T) {
	types := map[string]string{"id": "integer", "score": "number", "caption": "string"}

	if !itemParamIsNumber(stmlparser.ParamBind{Source: "item.id"}, types) {
		t.Error("integer item field should be numeric")
	}
	if !itemParamIsNumber(stmlparser.ParamBind{Source: "item.score"}, types) {
		t.Error("number item field should be numeric")
	}
	if itemParamIsNumber(stmlparser.ParamBind{Source: "item.caption"}, types) {
		t.Error("string item field is not numeric")
	}
	if itemParamIsNumber(stmlparser.ParamBind{Source: "item.unknown"}, types) {
		t.Error("unknown item field is not numeric")
	}
	if itemParamIsNumber(stmlparser.ParamBind{Source: "route.BuildingID"}, types) {
		t.Error("non-item source is never numeric here")
	}
	// dotted path resolves the first segment
	dotted := map[string]string{"photo": "integer"}
	if !itemParamIsNumber(stmlparser.ParamBind{Source: "item.photo.id"}, dotted) {
		t.Error("dotted path should resolve its first segment")
	}
	if itemParamIsNumber(stmlparser.ParamBind{Source: "item.id"}, nil) {
		t.Error("nil type map yields false")
	}
}

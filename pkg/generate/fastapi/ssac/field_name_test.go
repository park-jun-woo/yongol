//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderArgValueLegacy — Location 미설정 시 source 기반 매핑 분기 전체 커버
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestFieldName(t *testing.T) {
	if got := fieldName(ir.FieldArg{Field: ".ID"}); got != "ID" {
		t.Errorf("got %q, want ID", got)
	}
	if got := fieldName(ir.FieldArg{Field: "ID"}); got != "ID" {
		t.Errorf("got %q, want ID", got)
	}
	if got := fieldName(ir.FieldArg{}); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}

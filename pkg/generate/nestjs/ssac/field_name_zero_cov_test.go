//ff:func feature=gen-nestjs type=test control=sequence
//ff:what nestjs/ssac 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestFieldName_ZeroCov(t *testing.T) {
	if got := fieldName(ir.FieldArg{Field: ".ID"}); got != "ID" {
		t.Errorf("dotted: %q", got)
	}
	if got := fieldName(ir.FieldArg{Field: "Name"}); got != "Name" {
		t.Errorf("plain: %q", got)
	}
}

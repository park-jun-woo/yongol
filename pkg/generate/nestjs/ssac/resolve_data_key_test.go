//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestResolveDataKey(t *testing.T) {
	if got := resolveDataKey(ir.FieldArg{ColumnName: "user_id"}); got != "user_id" {
		t.Errorf("columnname = %q", got)
	}
	if got := resolveDataKey(ir.FieldArg{Key: "UserName"}); got != toSnake("UserName") {
		t.Errorf("key = %q", got)
	}
	if got := resolveDataKey(ir.FieldArg{Field: ".CourseId"}); got != toSnake("CourseId") {
		t.Errorf("field = %q", got)
	}
	if got := resolveDataKey(ir.FieldArg{}); got != "" {
		t.Errorf("empty = %q", got)
	}
}

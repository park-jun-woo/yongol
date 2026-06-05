//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderQueryParams — renderQueryParams 메타 목록→Python 파라미터 선언 목록 검증
package ssac

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderQueryParams(t *testing.T) {
	in := []ir.QueryParamMeta{
		{Name: "limit", Type: "integer", Required: true},
		{Name: "cursor", Type: "string", Required: false},
	}
	got := renderQueryParams(in)
	want := []string{"limit: int", "cursor: str | None = None"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("renderQueryParams() = %v, want %v", got, want)
	}

	if empty := renderQueryParams(nil); len(empty) != 0 {
		t.Errorf("renderQueryParams(nil) = %v, want empty", empty)
	}
}

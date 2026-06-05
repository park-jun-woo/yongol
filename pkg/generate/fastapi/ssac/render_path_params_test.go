//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderPathParams — renderPathParams path 파라미터→int 타입 선언 목록 검증
package ssac

import (
	"reflect"
	"testing"
)

func TestRenderPathParams(t *testing.T) {
	got := renderPathParams([]string{"id", "order_id"})
	want := []string{"id: int", "order_id: int"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("renderPathParams() = %v, want %v", got, want)
	}

	if empty := renderPathParams(nil); len(empty) != 0 {
		t.Errorf("renderPathParams(nil) = %v, want empty", empty)
	}
}

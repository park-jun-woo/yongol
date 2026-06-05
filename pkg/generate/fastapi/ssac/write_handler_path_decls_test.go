//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteHandlerPathDecls — writeHandlerPathDecls path 파라미터 int 선언 출력 검증
package ssac

import (
	"strings"
	"testing"
)

func TestWriteHandlerPathDecls(t *testing.T) {
	var b strings.Builder
	writeHandlerPathDecls(&b, []string{"id", "order_id"})
	want := "    id: int,\n    order_id: int,\n"
	if got := b.String(); got != want {
		t.Errorf("writeHandlerPathDecls() = %q, want %q", got, want)
	}

	var empty strings.Builder
	writeHandlerPathDecls(&empty, nil)
	if empty.String() != "" {
		t.Errorf("nil path params should write nothing, got %q", empty.String())
	}
}

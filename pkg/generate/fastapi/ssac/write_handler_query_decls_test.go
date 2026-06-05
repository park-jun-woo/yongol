//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteHandlerQueryDecls — writeHandlerQueryDecls query 파라미터 선언 출력 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteHandlerQueryDecls(t *testing.T) {
	var b strings.Builder
	writeHandlerQueryDecls(&b, []ir.QueryParamMeta{
		{Name: "limit", Type: "integer", Required: true},
		{Name: "cursor", Type: "string", Required: false},
	})
	want := "    limit: int,\n    cursor: str | None = None,\n"
	if got := b.String(); got != want {
		t.Errorf("writeHandlerQueryDecls() = %q, want %q", got, want)
	}

	var empty strings.Builder
	writeHandlerQueryDecls(&empty, nil)
	if empty.String() != "" {
		t.Errorf("nil query params should write nothing, got %q", empty.String())
	}
}

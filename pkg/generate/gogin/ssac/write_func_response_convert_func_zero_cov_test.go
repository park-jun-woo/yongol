//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestWriteFuncResponseConvertFunc_ZeroCov — Func 응답 → api 변환 (required value / optional ptr)
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestWriteFuncResponseConvertFunc_ZeroCov(t *testing.T) {
	var sb strings.Builder
	info := funcRespInfo{PkgAlias: "dashboard", ImportPath: "example.com/app/internal/dashboard"}
	spec := &funcspec.FuncSpec{
		Name: "Summarize",
		ResponseFields: []funcspec.Field{
			{Name: "Total", JSONName: "total"},
		},
	}
	writeFuncResponseConvertFunc(&sb, "SummarizeResponse", convertSchemaZeroCov(), info, spec)
	out := sb.String()
	for _, want := range []string{
		"func convertSummarizeResponse(src dashboard.SummarizeResponse) (*api.SummarizeResponse, error) {",
		"return &api.SummarizeResponse{",
		"Id: src.Id,",      // required → value
		"Name: &src.Name,", // optional → pointer
		"}, nil",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in:\n%s", want, out)
		}
	}
}

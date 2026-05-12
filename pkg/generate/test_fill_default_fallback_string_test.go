//ff:func feature=generate type=test control=sequence
//ff:what OpenAPI에 없는 operationId가 string 타입으로 폴백되는지 검증

package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFillDefaultRequestConstraints_FallbackToString(t *testing.T) {
	// Operation not in OpenAPI doc at all — fallback to string.
	doc := loadTestDoc(t)
	pages := []stmlparser.PageSpec{
		{
			Actions: []stmlparser.ActionBlock{
				{
					OperationID: "NonExistent",
					Fields:      []stmlparser.FieldBind{{Name: "data"}},
				},
			},
		},
	}
	result := fillDefaultRequestConstraints(pages, doc, nil)
	fields, ok := result["NonExistent"]
	if !ok {
		t.Fatal("NonExistent not found in result")
	}
	if got := fields["data"].Type; got != "string" {
		t.Errorf("data.Type = %q, want %q", got, "string")
	}
}

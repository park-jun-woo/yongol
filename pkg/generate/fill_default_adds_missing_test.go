//ff:func feature=generate type=test control=sequence
//ff:what fillDefaultRequestConstraints가 누락된 operationId를 기본 타입으로 채우는지 검증

package generate

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFillDefaultRequestConstraints_AddsMissing(t *testing.T) {
	doc := loadTestDoc(t)
	pages := []stmlparser.PageSpec{
		{
			Actions: []stmlparser.ActionBlock{
				{
					OperationID: "CreateItem",
					Fields: []stmlparser.FieldBind{
						{Name: "title"},
						{Name: "count"},
					},
				},
			},
		},
	}
	existing := map[string]map[string]oapiparser.FieldConstraint{}
	result := fillDefaultRequestConstraints(pages, doc, existing)

	fields, ok := result["CreateItem"]
	if !ok {
		t.Fatal("CreateItem not found in result")
	}
	if got := fields["title"].Type; got != "string" {
		t.Errorf("title.Type = %q, want %q", got, "string")
	}
	if got := fields["title"].Required; !got {
		t.Error("title.Required = false, want true")
	}
	if got := fields["count"].Type; got != "integer" {
		t.Errorf("count.Type = %q, want %q", got, "integer")
	}
}

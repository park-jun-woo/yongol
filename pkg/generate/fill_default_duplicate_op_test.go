//ff:func feature=generate type=test control=sequence
//ff:what fillDefaultRequestConstraints가 중복 operationId의 두 번째 항목을 건너뛰는지 검증

package generate

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFillDefaultRequestConstraints_DuplicateOp(t *testing.T) {
	doc := loadTestDoc(t)
	// Two form actions with the same OperationID -> first resolves, second
	// hits the `result[ae.opID] done` skip branch.
	pages := []stmlparser.PageSpec{
		{
			Actions: []stmlparser.ActionBlock{
				{
					OperationID: "CreateItem",
					Fields:      []stmlparser.FieldBind{{Name: "title"}},
				},
			},
		},
		{
			Actions: []stmlparser.ActionBlock{
				{
					OperationID: "CreateItem",
					Fields:      []stmlparser.FieldBind{{Name: "title"}},
				},
			},
		},
	}
	existing := map[string]map[string]oapiparser.FieldConstraint{}
	result := fillDefaultRequestConstraints(pages, doc, existing)

	if _, ok := result["CreateItem"]; !ok {
		t.Fatal("CreateItem not found in result")
	}
	if result["CreateItem"]["title"].Type != "string" {
		t.Errorf("title.Type = %q, want string", result["CreateItem"]["title"].Type)
	}
}

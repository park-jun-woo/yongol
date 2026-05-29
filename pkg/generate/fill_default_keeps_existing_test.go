//ff:func feature=generate type=test control=sequence
//ff:what fillDefaultRequestConstraints가 기존 constraint를 덮어쓰지 않는지 검증

package generate

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFillDefaultRequestConstraints_KeepsExisting(t *testing.T) {
	doc := loadTestDoc(t)
	maxLen := 100
	existing := map[string]map[string]oapiparser.FieldConstraint{
		"CreateItem": {
			"title": {Type: "string", Required: true, MaxLength: &maxLen},
		},
	}
	pages := []stmlparser.PageSpec{
		{
			Actions: []stmlparser.ActionBlock{
				{
					OperationID: "CreateItem",
					Fields:      []stmlparser.FieldBind{{Name: "title"}},
				},
			},
		},
	}
	result := fillDefaultRequestConstraints(pages, doc, existing)
	fc := result["CreateItem"]["title"]
	if fc.MaxLength == nil || *fc.MaxLength != 100 {
		t.Error("existing constraint was overwritten")
	}
}

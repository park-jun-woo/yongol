//ff:func feature=generate type=test control=sequence
//ff:what fillDefaultRequestConstraints의 early-return 분기(needed 없음/missing 없음) 검증

package generate

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFillDefaultRequestConstraints_NoFormActions(t *testing.T) {
	doc := loadTestDoc(t)
	// Pages with no form fields -> needed is empty -> existing returned as-is.
	pages := []stmlparser.PageSpec{{}}
	existing := map[string]map[string]oapiparser.FieldConstraint{
		"Untouched": {"x": {Type: "string"}},
	}
	result := fillDefaultRequestConstraints(pages, doc, existing)
	if len(result) != 1 || result["Untouched"]["x"].Type != "string" {
		t.Errorf("expected existing returned unchanged, got: %v", result)
	}
}

func TestFillDefaultRequestConstraints_AlreadyCovered(t *testing.T) {
	doc := loadTestDoc(t)
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
	// CreateItem already has non-empty constraints -> nothing missing.
	existing := map[string]map[string]oapiparser.FieldConstraint{
		"CreateItem": {"title": {Type: "string", Required: true}},
	}
	result := fillDefaultRequestConstraints(pages, doc, existing)
	// Same map returned (no augmentation needed).
	if len(result) != 1 {
		t.Errorf("expected no new entries, got: %v", result)
	}
	if result["CreateItem"]["title"].Type != "string" {
		t.Errorf("existing constraint should be preserved, got: %v", result["CreateItem"])
	}
}

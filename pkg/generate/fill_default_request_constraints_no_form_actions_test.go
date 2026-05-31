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

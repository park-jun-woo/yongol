//ff:func feature=generate type=test control=sequence
//ff:what 액션이 없을 때 fillDefaultRequestConstraints가 기존 맵을 그대로 반환하는지 검증

package generate

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

func TestFillDefaultRequestConstraints_NoActions(t *testing.T) {
	existing := map[string]map[string]oapiparser.FieldConstraint{
		"Op": {"f": {Type: "string"}},
	}
	result := fillDefaultRequestConstraints(nil, nil, existing)
	if len(result) != len(existing) {
		t.Error("result should equal existing when no pages")
	}
}

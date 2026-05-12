//ff:func feature=generate type=test control=sequence
//ff:what 필드가 있는 액션이 field-less 수집에서 제외되는지 검증

package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectNestedFieldlessActions_SkipsWithFields(t *testing.T) {
	ab := stmlparser.ActionBlock{
		OperationID: "UpdateBuilding",
		Fields:      []stmlparser.FieldBind{{Name: "name"}},
	}
	nodes := []stmlparser.ChildNode{
		{Kind: "action", Action: &ab},
	}
	result := make(map[string]bool)
	collectNestedFieldlessActions(nodes, result)
	if result["UpdateBuilding"] {
		t.Error("UpdateBuilding has fields, should not be collected")
	}
}

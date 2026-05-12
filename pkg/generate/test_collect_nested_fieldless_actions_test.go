//ff:func feature=generate type=test control=sequence
//ff:what each 블록 내 field-less 중첩 액션이 수집되는지 검증

package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectNestedFieldlessActions_Each(t *testing.T) {
	ab := stmlparser.ActionBlock{OperationID: "DeleteBuilding", Fields: nil}
	nodes := []stmlparser.ChildNode{
		{
			Kind: "each",
			Each: &stmlparser.EachBlock{
				Field: "buildings",
				Children: []stmlparser.ChildNode{
					{Kind: "action", Action: &ab},
				},
			},
		},
	}
	result := make(map[string]bool)
	collectNestedFieldlessActions(nodes, result)
	if !result["DeleteBuilding"] {
		t.Error("expected DeleteBuilding to be collected from each block")
	}
}

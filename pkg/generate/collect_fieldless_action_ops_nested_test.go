//ff:func feature=generate type=test control=sequence
//ff:what STML 중첩 field-less 액션이 NoBodyOps 후보로 수집되는지 검증

package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectFieldlessActionOps_Nested(t *testing.T) {
	ab := stmlparser.ActionBlock{OperationID: "CancelReservation", Fields: nil}
	pages := []stmlparser.PageSpec{
		{
			Name: "test-page",
			Children: []stmlparser.ChildNode{
				{
					Kind: "fetch",
					Fetch: &stmlparser.FetchBlock{
						OperationID: "ListReservations",
						Children: []stmlparser.ChildNode{
							{
								Kind: "state",
								State: &stmlparser.StateBind{
									Condition: "canCancel",
									Children: []stmlparser.ChildNode{
										{Kind: "action", Action: &ab},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	result := collectFieldlessActionOps(pages)
	if !result["CancelReservation"] {
		t.Error("expected CancelReservation to be collected from nested state")
	}
}

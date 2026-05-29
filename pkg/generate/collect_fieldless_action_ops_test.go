//ff:func feature=generate type=test control=sequence
//ff:what STML top-level field-less 액션이 NoBodyOps 후보로 수집되는지 검증

package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectFieldlessActionOps_TopLevel(t *testing.T) {
	pages := []stmlparser.PageSpec{
		{
			Name: "test-page",
			Actions: []stmlparser.ActionBlock{
				{OperationID: "DeleteRoom", Fields: nil},
				{OperationID: "CreateRoom", Fields: []stmlparser.FieldBind{{Name: "name"}}},
			},
		},
	}
	result := collectFieldlessActionOps(pages)
	if !result["DeleteRoom"] {
		t.Error("expected DeleteRoom to be field-less")
	}
	if result["CreateRoom"] {
		t.Error("CreateRoom has fields, should not be in result")
	}
}

//ff:func feature=generate type=test control=sequence
//ff:what TestByName_ZeroCov — generate 폼 액션/필드 해석 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package generate

import (
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func byNameFormPage() stmlparser.PageSpec {
	form := stmlparser.ActionBlock{
		OperationID: "CreateItem",
		Fields:      []stmlparser.FieldBind{{Name: "Name"}, {Name: "Count"}},
	}
	fieldless := stmlparser.ActionBlock{OperationID: "DeleteItem"}
	nestedForm := stmlparser.ActionBlock{
		OperationID: "UpdateItem",
		Fields:      []stmlparser.FieldBind{{Name: "Title"}},
	}
	return stmlparser.PageSpec{
		FileName: "items.html",
		Actions:  []stmlparser.ActionBlock{form, fieldless},
		Children: []stmlparser.ChildNode{
			{Kind: "action", Action: &nestedForm},
			{Kind: "fetch", Fetch: &stmlparser.FetchBlock{
				OperationID: "ListItems",
				Children: []stmlparser.ChildNode{
					{Kind: "action", Action: &form},
				},
			}},
		},
	}
}

//ff:func feature=stml-gen type=test control=sequence
//ff:what setRowActionArgsInFetch — flat Eaches·Children·NestedFetches 경유 RowMutateArg 설정 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetRowActionArgsInFetch(t *testing.T) {
	itemTypes := map[string]map[string]map[string]string{
		"ListPhotos": {"photos": {"id": "integer"}},
		"ListTags":   {"tags": {"name": "string"}},
	}

	flatAction := stmlparser.ActionBlock{
		OperationID: "DeletePhoto",
		Params:      []stmlparser.ParamBind{{Name: "photoId", Source: "item.id"}},
	}
	flatEach := stmlparser.EachBlock{Field: "photos", Children: []stmlparser.ChildNode{
		{Kind: "action", Action: &flatAction},
	}}

	nestedAction := stmlparser.ActionBlock{
		OperationID: "DropTag",
		Params:      []stmlparser.ParamBind{{Name: "tagName", Source: "item.name"}},
	}
	nestedEach := stmlparser.EachBlock{Field: "tags", Children: []stmlparser.ChildNode{
		{Kind: "action", Action: &nestedAction},
	}}
	nested := stmlparser.FetchBlock{
		OperationID: "ListTags",
		Children:    []stmlparser.ChildNode{{Kind: "each", Each: &nestedEach}},
	}

	f := stmlparser.FetchBlock{
		OperationID:   "ListPhotos",
		Eaches:        []stmlparser.EachBlock{flatEach},
		NestedFetches: []stmlparser.FetchBlock{nested},
	}
	setRowActionArgsInFetch(&f, "ListPhotos", itemTypes, nil)

	if flatAction.RowMutateArg != "{ photoId: item.id }" {
		t.Errorf("flat each action: %q", flatAction.RowMutateArg)
	}
	// the nested fetch resolves its own operationId for the item schema
	if nestedAction.RowMutateArg != "{ tagName: item.name }" {
		t.Errorf("nested fetch action: %q", nestedAction.RowMutateArg)
	}
}

//ff:func feature=stml-gen type=test control=sequence
//ff:what renderRowMutateArg — 인자 객체 리터럴 생성(정수 래핑·숫자 item 생략·빈 파라미터) 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderRowMutateArg(t *testing.T) {
	a := stmlparser.ActionBlock{
		OperationID: "DeletePhoto",
		Params: []stmlparser.ParamBind{
			{Name: "buildingId", Source: "route.BuildingID"},
			{Name: "photoId", Source: "item.id"},
			{Name: "tag", Source: "item.tag"},
		},
	}
	ppt := map[string]map[string]string{
		"DeletePhoto": {"buildingId": "integer", "photoId": "integer", "tag": "string"},
	}

	// item.id is integer in the item schema → no wrapping; route param and
	// the string-typed integer path param keep their rules.
	got := renderRowMutateArg(a, ppt, map[string]string{"id": "integer", "tag": "string"})
	want := "{ buildingId: Number(BuildingID), photoId: item.id, tag: item.tag }"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	// unknown item schema → integer path params are Number-wrapped.
	got = renderRowMutateArg(a, ppt, nil)
	want = "{ buildingId: Number(BuildingID), photoId: Number(item.id), tag: item.tag }"
	if got != want {
		t.Errorf("nil schema: got %q, want %q", got, want)
	}

	// no params → empty string (caller keeps the default mutate arg).
	if got := renderRowMutateArg(stmlparser.ActionBlock{}, nil, nil); got != "" {
		t.Errorf("no params: got %q, want empty", got)
	}
}

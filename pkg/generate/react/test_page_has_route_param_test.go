//ff:func feature=gen-react type=test control=sequence
//ff:what pageHasRouteParam route 파라미터 존재 여부 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestPageHasRouteParam(t *testing.T) {
	noRoute := stml.PageSpec{
		Name:     "login",
		FileName: "login.html",
	}
	if pageHasRouteParam(noRoute) {
		t.Error("expected no route param for login page")
	}

	withFetchRoute := stml.PageSpec{
		Name:     "workflow-detail",
		FileName: "workflow-detail.html",
		Fetches: []stml.FetchBlock{{
			Params: []stml.ParamBind{{Name: "id", Source: "route.id"}},
		}},
	}
	if !pageHasRouteParam(withFetchRoute) {
		t.Error("expected route param for page with route.id in fetch")
	}

	withActionRoute := stml.PageSpec{
		Name:     "room-edit",
		FileName: "room-edit.html",
		Actions: []stml.ActionBlock{{
			Params: []stml.ParamBind{{Name: "RoomID", Source: "route.RoomID"}},
		}},
	}
	if !pageHasRouteParam(withActionRoute) {
		t.Error("expected route param for page with route.RoomID in action")
	}
}

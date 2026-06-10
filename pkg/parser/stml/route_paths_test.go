//ff:func feature=stml-parse type=test control=sequence
//ff:what TestRoutePaths — 명시 라우트·detail·route param·기본 파일명 유래 패턴 검증

package stml

import (
	"reflect"
	"testing"
)

func TestRoutePaths(t *testing.T) {
	// Explicit data-route wins as-is.
	got := RoutePaths(PageSpec{FileName: "login.html", Route: "/sign-in"})
	if !reflect.DeepEqual(got, []string{"/sign-in"}) {
		t.Errorf("explicit: got %v", got)
	}

	// "-detail" suffix → pluralized parent + /:id.
	got = RoutePaths(PageSpec{FileName: "workflow-detail.html"})
	if !reflect.DeepEqual(got, []string{"/workflows/:id"}) {
		t.Errorf("detail: got %v", got)
	}

	// Plain page → filename-derived path.
	got = RoutePaths(PageSpec{FileName: "login.html"})
	if !reflect.DeepEqual(got, []string{"/login"}) {
		t.Errorf("plain: got %v", got)
	}

	// Route param → base + base/:id.
	p := PageSpec{
		FileName: "templates.html",
		Fetches:  []FetchBlock{{OperationID: "GetTemplate", Params: []ParamBind{{Name: "id", Source: "route.id"}}}},
	}
	got = RoutePaths(p)
	if !reflect.DeepEqual(got, []string{"/templates", "/templates/:id"}) {
		t.Errorf("route param: got %v", got)
	}
}

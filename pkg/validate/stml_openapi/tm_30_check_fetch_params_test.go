//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm30CheckFetchParams — fetch 파라미터의 item.* 발화와 route.* 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM30CheckFetchParams(t *testing.T) {
	f := stml.FetchBlock{
		OperationID: "GetUnit",
		Params: []stml.ParamBind{
			{Name: "buildingId", Source: "route.BuildingID"},
			{Name: "photoId", Source: "item.id"},
		},
	}
	got := tm30CheckFetchParams(f, "p.html")
	if len(got) != 1 {
		t.Fatalf("expected 1 diag, got %v", got)
	}
	if !strings.Contains(got[0].Message, "[TM-30]") || got[0].OperationID != "GetUnit" {
		t.Errorf("unexpected diag: %+v", got[0])
	}

	clean := stml.FetchBlock{
		OperationID: "GetUnit",
		Params:      []stml.ParamBind{{Name: "buildingId", Source: "route.BuildingID"}},
	}
	if got := tm30CheckFetchParams(clean, "p.html"); len(got) != 0 {
		t.Errorf("route-only fetch must stay silent: %v", got)
	}
}

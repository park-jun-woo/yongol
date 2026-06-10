//ff:func feature=stml-gen type=test control=sequence
//ff:what renderUseParamsWithRoute optional 세그먼트(:Name?) 패턴의 ? 마커가 strip 되는지 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderUseParamsWithRoute_OptionalPatternStripsMarker(t *testing.T) {
	got := renderUseParamsWithRoute(nil, "/unit-info/:BuildingID/:UnitID/:PhotoID?")
	want := "const { BuildingID, UnitID, PhotoID } = useParams()"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	// bind names and optional pattern params merge without duplicates
	binds := []stmlparser.ParamBind{{Name: "buildingId", Source: "route.BuildingID"}}
	got = renderUseParamsWithRoute(binds, "/unit-info/:BuildingID/:PhotoID?")
	want = "const { BuildingID, PhotoID } = useParams()"
	if got != want {
		t.Errorf("merged: got %q, want %q", got, want)
	}
}

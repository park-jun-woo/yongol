//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what data-route 패턴에서 파라미터 이름 추출 테이블 테스트

package stml

import (
	"slices"
	"testing"
)

func TestExtractRoutePatternParams(t *testing.T) {
	tests := []struct {
		route string
		want  []string
	}{
		{"", nil},
		{"/buildings", nil},
		{"/buildings/:buildingId/units/:id", []string{"buildingId", "id"}},
		{"/workflows/:id", []string{"id"}},
		{"/a/:x/b/:y/c/:z", []string{"x", "y", "z"}},
		// optional segment markers are stripped so useParams gets exact names
		{"/unit-info/:BuildingID/:UnitID/:PhotoID?", []string{"BuildingID", "UnitID", "PhotoID"}},
		{"/webhooks/:id?", []string{"id"}},
	}
	for _, tt := range tests {
		got := extractRoutePatternParams(tt.route)
		if !slices.Equal(got, tt.want) {
			t.Errorf("extractRoutePatternParams(%q) = %v, want %v", tt.route, got, tt.want)
		}
	}
}

//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what routeSegmentName — route.<Name>/비route/빈 세그먼트 소스 파싱 검증

package stml

import "testing"

func TestRouteSegmentName(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		wantName string
		wantOK   bool
	}{
		{name: "valid route source", source: "route.BuildingID", wantName: "BuildingID", wantOK: true},
		{name: "non-route source", source: "query.q", wantName: "", wantOK: false},
		{name: "empty source", source: "", wantName: "", wantOK: false},
		{name: "route prefix with empty segment", source: "route.", wantName: "", wantOK: false},
		{name: "item source", source: "item.id", wantName: "", wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			name, ok := routeSegmentName(tt.source)
			if name != tt.wantName || ok != tt.wantOK {
				t.Errorf("routeSegmentName(%q) = (%q, %v), want (%q, %v)", tt.source, name, ok, tt.wantName, tt.wantOK)
			}
		})
	}
}

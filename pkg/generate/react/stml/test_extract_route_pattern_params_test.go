//ff:func feature=stml-gen type=test control=sequence
//ff:what data-route 패턴에서 파라미터 이름 추출 테스트

package stml

import (
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
	}
	for _, tt := range tests {
		got := extractRoutePatternParams(tt.route)
		if len(got) != len(tt.want) {
			t.Errorf("extractRoutePatternParams(%q) = %v, want %v", tt.route, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("extractRoutePatternParams(%q)[%d] = %q, want %q", tt.route, i, got[i], tt.want[i])
			}
		}
	}
}

func TestMergeRouteParamNames(t *testing.T) {
	tests := []struct {
		a, b []string
		want []string
	}{
		{nil, nil, nil},
		{[]string{"id"}, nil, []string{"id"}},
		{nil, []string{"id"}, []string{"id"}},
		{[]string{"id"}, []string{"id"}, []string{"id"}},
		{[]string{"buildingId"}, []string{"buildingId", "id"}, []string{"buildingId", "id"}},
		{[]string{"id"}, []string{"buildingId"}, []string{"id", "buildingId"}},
	}
	for _, tt := range tests {
		got := mergeRouteParamNames(tt.a, tt.b)
		if len(got) != len(tt.want) {
			t.Errorf("mergeRouteParamNames(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			continue
		}
		for i := range got {
			if got[i] != tt.want[i] {
				t.Errorf("mergeRouteParamNames(%v, %v)[%d] = %q, want %q", tt.a, tt.b, i, got[i], tt.want[i])
			}
		}
	}
}

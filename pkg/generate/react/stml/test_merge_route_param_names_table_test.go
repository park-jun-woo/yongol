//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what mergeRouteParamNames 병합 테이블 테스트

package stml

import (
	"slices"
	"testing"
)

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
		if !slices.Equal(got, tt.want) {
			t.Errorf("mergeRouteParamNames(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
		}
	}
}

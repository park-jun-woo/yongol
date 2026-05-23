//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-openapi
//ff:what sortedKeys — nil/empty/복수 키 오름차순 정렬 검증

package openapi_ssac

import (
	"reflect"
	"testing"
)

func TestSortedKeys(t *testing.T) {
	tests := []struct {
		name string
		m    map[int]bool
		want []int
	}{
		{"nil", nil, []int{}},
		{"empty", map[int]bool{}, []int{}},
		{"single", map[int]bool{200: true}, []int{200}},
		{"multiple sorted", map[int]bool{500: true, 200: true, 404: true}, []int{200, 404, 500}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedKeys(tt.m)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

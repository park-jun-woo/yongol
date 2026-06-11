//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what addBreadcrumbPageEdges — 조상 간선 추가 / 자기 자신 제외 / 기존 간선 중복 억제 검증

package stml_openapi

import (
	"reflect"
	"testing"
)

func TestAddBreadcrumbPageEdges(t *testing.T) {
	g := &pageGraph{Edges: map[string][]string{
		"building-detail": {"building-list"}, // pre-existing redirect edge
	}}

	addBreadcrumbPageEdges("building-detail", []string{"dashboard", "building-list", "building-detail"}, g)

	want := []string{"building-list", "dashboard"}
	if !reflect.DeepEqual(g.Edges["building-detail"], want) {
		t.Errorf("Edges[building-detail] = %v, want %v (dedupe + self excluded)", g.Edges["building-detail"], want)
	}
}

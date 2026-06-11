//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what appendEdgeOnce — 신규 간선 추가 / 동일 간선 중복 억제 / 다른 대상은 누적 검증

package stml_openapi

import (
	"reflect"
	"testing"
)

func TestAppendEdgeOnce(t *testing.T) {
	g := &pageGraph{Edges: map[string][]string{}}

	appendEdgeOnce(g, "building-detail", "building-list")
	appendEdgeOnce(g, "building-detail", "building-list") // duplicate — suppressed
	appendEdgeOnce(g, "building-detail", "dashboard")

	want := []string{"building-list", "dashboard"}
	if !reflect.DeepEqual(g.Edges["building-detail"], want) {
		t.Errorf("Edges[building-detail] = %v, want %v", g.Edges["building-detail"], want)
	}
}

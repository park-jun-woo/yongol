//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestRunDomained_FrontendOffSkipsXMO12 — 도메인 모드 frontend OFF 시 XMO-12 집계 단락 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestRunDomained_FrontendOffSkipsXMO12(t *testing.T) {
	docs := map[string]*openapi3.T{
		"public": makeDoc(map[string]*openapi3.PathItem{
			"/users": {Get: &openapi3.Operation{OperationID: "ListUsers", Tags: []string{noFrontTag}}},
		}),
	}
	fs := domainedFS(docs, nil)
	// Frontend OFF (no Lang/Framework) → xmo12 aggregate must short-circuit.
	fs.Manifest.Frontend = manifest.Frontend{}
	if got := xmo12NoFrontConsumedAll(fs); got != nil {
		t.Errorf("frontend off: expected nil, got %v", got)
	}
}

//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what zz_zerocov_test — ssac_sqlc 0% 헬퍼 (Run / collectInputKeys / buildQueryParamMap / checkSingleInputKeyCase / checkSeqInputKeyCase) 단위 테스트
package ssac_sqlc

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildXqs18OperationMap_ZeroCov(t *testing.T) {
	// nil doc → empty map.
	if m := buildXqs18OperationMap(nil); len(m) != 0 {
		t.Errorf("nil doc should give empty map, got %v", m)
	}

	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/users", &openapi3.PathItem{
		Get:  &openapi3.Operation{OperationID: "ListUsers"},
		Post: &openapi3.Operation{OperationID: ""}, // skipped: no opID
	})
	m := buildXqs18OperationMap(doc)
	if m["ListUsers"] == nil {
		t.Errorf("expected ListUsers in map, got %v", m)
	}
	if len(m) != 1 {
		t.Errorf("expected exactly 1 op, got %d", len(m))
	}
}

//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what addFetchOps — 단일 fetch와 중첩 fetch의 operationId를 집합에 재귀 추가 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAddFetchOps(t *testing.T) {
	ops := map[string]bool{}
	addFetchOps(stml.FetchBlock{
		OperationID: "A",
		NestedFetches: []stml.FetchBlock{
			{OperationID: "B", NestedFetches: []stml.FetchBlock{{OperationID: "C"}}},
		},
	}, ops)
	for _, want := range []string{"A", "B", "C"} {
		if !ops[want] {
			t.Errorf("missing %q in %+v", want, ops)
		}
	}
}

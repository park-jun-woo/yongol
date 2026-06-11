//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestCollectConsumedOpsWithComponents — fetch + 컴포넌트 .tsx api 호출 합산 소비

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectConsumedOpsWithComponents(t *testing.T) {
	specsDir := t.TempDir()
	writeComponent(t, specsDir, "ChildComp", `api.CompOp();`)

	pages := []stml.PageSpec{{
		Name:    "page",
		Fetches: []stml.FetchBlock{{OperationID: "ListItems"}},
		Children: []stml.ChildNode{{
			Kind:      "component",
			Component: &stml.ComponentRef{Name: "ChildComp"},
		}},
	}}
	ops := map[string]struct{}{"CompOp": {}}

	out := collectConsumedOps(pages, nil, nil, specsDir, ops)
	for _, want := range []string{"ListItems", "CompOp"} {
		if _, ok := out[want]; !ok {
			t.Errorf("missing consumed op %q", want)
		}
	}
}

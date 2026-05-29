//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_Unconsumed_NestedFetchConsumed

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_Unconsumed_NestedFetchConsumed(t *testing.T) {
	// Nested fetch's operationId should be counted as consumed.
	pages := []stml.PageSpec{{
		FileName: "detail.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetWorkflow",
			NestedFetches: []stml.FetchBlock{{
				OperationID: "ListWorkflowVersions",
			}},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/workflows/{id}":          getOp("GetWorkflow", nil, nil),
		"/workflows/{id}/versions": getOp("ListWorkflowVersions", nil, nil),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("unexpected XMO-10 for nested fetch, got %v", diags)
	}
}

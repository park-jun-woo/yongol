//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM01_FetchOpNotFound_Positive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM01_FetchOpNotFound_Positive(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "dashboard.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "NoSuchOp",
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, nil),
	})
	diags := Run(makeFS(pages, doc))
	if !hasDiag(diags, "[TM-01]") {
		t.Errorf("expected TM-01 diagnostic, got %v", diags)
	}
}

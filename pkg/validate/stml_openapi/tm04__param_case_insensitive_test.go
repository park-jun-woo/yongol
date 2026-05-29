//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM04_ParamCaseInsensitive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM04_ParamCaseInsensitive(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "detail.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetUser",
			Params:      []stml.ParamBind{{Name: "ID", Source: "route.id"}},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users/{id}": getOp("GetUser", []*openapi3.ParameterRef{
			paramRef("id", "path"),
		}, nil),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[TM-04]") {
		t.Errorf("case-insensitive match should pass, got %v", diags)
	}
}

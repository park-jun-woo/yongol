//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-08 test — data-each 필드가 array 타입이 아닌 경우 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM08_EachNotArray_Positive(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "list.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
			Eaches: []stml.EachBlock{
				{Field: "totalCount"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, map[string]*openapi3.SchemaRef{
			"totalCount": intProp(),
			"users":      arrayProp("object"),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if !hasDiag(diags, "[TM-08]") {
		t.Errorf("expected TM-08 diagnostic, got %v", diags)
	}
}

func TestTM08_EachNotArray_Negative(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "list.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
			Eaches: []stml.EachBlock{
				{Field: "users"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/users": getOp("ListUsers", nil, map[string]*openapi3.SchemaRef{
			"totalCount": intProp(),
			"users":      arrayProp("object"),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if hasDiag(diags, "[TM-08]") {
		t.Errorf("unexpected TM-08 diagnostic, got %v", diags)
	}
}

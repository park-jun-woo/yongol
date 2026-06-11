//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM06_BindNotFound_Positive

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM06_BindNotFound_Positive(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "profile.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetProfile",
			Binds: []stml.FieldBind{
				{Name: "NonExistent"},
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/profile": getOp("GetProfile", nil, map[string]*openapi3.SchemaRef{
			"Name":  stringProp(),
			"Email": stringProp(),
		}),
	})
	diags := Run(makeFS(pages, doc))
	if !hasDiag(diags, "[TM-06]") {
		t.Errorf("expected TM-06 diagnostic, got %v", diags)
	}

	// Phase039 / BUG-128: data-bind against a non-200 success body. A 201
	// (Created) op exposes its body fields, so binding an existing one passes;
	// a 204 (No Content) op has no body, so any bind must ERROR.
	item201 := postOpStatusResp("Create201", 201, map[string]*openapi3.SchemaRef{
		"Name": stringProp(),
	})
	if d := tm06Binds([]stml.FieldBind{{Name: "Name"}}, "Create201", "p.html", operationEntry{method: "POST", op: item201.Post}); len(d) != 0 {
		t.Errorf("201 body bind: expected 0 diagnostics, got %+v", d)
	}
	item204 := postOpStatusResp("Ping204", 204, nil)
	if d := tm06Binds([]stml.FieldBind{{Name: "Name"}}, "Ping204", "p.html", operationEntry{method: "POST", op: item204.Post}); len(d) != 1 {
		t.Errorf("204 bind: expected 1 TM-06, got %+v", d)
	}
}

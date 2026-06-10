//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what tm29ActionOnErrorMissing — 4xx/5xx 선언 op + data-on-error 부재 WARNING 발화/침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM29ActionOnErrorMissing(t *testing.T) {
	errResp := func() *openapi3.ResponseRef {
		desc := "error"
		return &openapi3.ResponseRef{Value: &openapi3.Response{Description: &desc}}
	}
	okResp := func() *openapi3.ResponseRef {
		desc := "ok"
		return &openapi3.ResponseRef{Value: &openapi3.Response{Description: &desc}}
	}
	// CreateItem: 200 + 400 + 500.
	create := &openapi3.Operation{
		OperationID: "CreateItem",
		Responses: openapi3.NewResponses(
			openapi3.WithStatus(200, okResp()),
			openapi3.WithStatus(400, errResp()),
			openapi3.WithStatus(500, errResp()),
		),
	}
	// DeleteItem: 5xx only.
	del := &openapi3.Operation{
		OperationID: "DeleteItem",
		Responses: openapi3.NewResponses(
			openapi3.WithStatus(204, okResp()),
			openapi3.WithStatus(503, errResp()),
		),
	}
	// UpdateItem: range wildcard "4XX".
	update := &openapi3.Operation{OperationID: "UpdateItem", Responses: openapi3.NewResponses(openapi3.WithStatus(200, okResp()))}
	update.Responses.Set("4XX", errResp())
	// CleanItem: 2xx only — no error contract declared.
	clean := &openapi3.Operation{
		OperationID: "CleanItem",
		Responses:   openapi3.NewResponses(openapi3.WithStatus(200, okResp())),
	}
	// DefaultItem: 2xx + "default" only — "default" is not a 4xx/5xx declaration.
	deflt := &openapi3.Operation{OperationID: "DefaultItem", Responses: openapi3.NewResponses(openapi3.WithStatus(200, okResp()))}
	deflt.Responses.Set("default", errResp())

	doc := makeDoc(map[string]*openapi3.PathItem{
		"/items":        {Post: create},
		"/items/del":    {Delete: del},
		"/items/update": {Put: update},
		"/items/clean":  {Post: clean},
		"/items/deflt":  {Post: deflt},
	})
	opMap := buildOperationMethodMap(doc)

	cases := []TestTM29ActionOnErrorMissingCase{
		// 4xx/5xx declared, no data-on-error → WARNING.
		{name: "error_responses_no_on_error", action: stml.ActionBlock{OperationID: "CreateItem"}, wantCount: 1},
		// 5xx-only declaration fires too.
		{name: "5xx_only_no_on_error", action: stml.ActionBlock{OperationID: "DeleteItem"}, wantCount: 1},
		// Range wildcard "4XX" counts as a declared error response.
		{name: "4XX_wildcard_no_on_error", action: stml.ActionBlock{OperationID: "UpdateItem"}, wantCount: 1},
		// data-on-error present → silent.
		{name: "error_responses_with_on_error", action: stml.ActionBlock{OperationID: "CreateItem", OnErrorNode: true}, wantCount: 0},
		// No 4xx/5xx declared → silent.
		{name: "no_error_responses", action: stml.ActionBlock{OperationID: "CleanItem"}, wantCount: 0},
		// "default" response alone is not an error declaration → silent.
		{name: "default_response_only", action: stml.ActionBlock{OperationID: "DefaultItem"}, wantCount: 0},
		// Unknown operationId → silent (TM-02 reports it).
		{name: "unknown_op_silent", action: stml.ActionBlock{OperationID: "Nope"}, wantCount: 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runTM29ActionOnErrorMissing(t, c, opMap)
		})
	}
}

//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what test helpers — STML↔OpenAPI 검증 테스트용 OpenAPI fixture 생성 유틸

package stml_openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// makeDoc builds an OpenAPI doc with the given paths.
func makeDoc(paths map[string]*openapi3.PathItem) *openapi3.T {
	p := &openapi3.Paths{}
	for path, item := range paths {
		p.Set(path, item)
	}
	return &openapi3.T{Paths: p}
}

// getOp creates a PathItem with a GET operation.
func getOp(opID string, params []*openapi3.ParameterRef, respProps map[string]*openapi3.SchemaRef) *openapi3.PathItem {
	op := &openapi3.Operation{OperationID: opID, Parameters: params}
	if respProps != nil {
		op.Responses = openapi3.NewResponses(
			openapi3.WithStatus(200, &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
						Type:       &openapi3.Types{"object"},
						Properties: respProps,
					}),
				},
			}),
		)
	}
	return &openapi3.PathItem{Get: op}
}

// postOp creates a PathItem with a POST operation.
func postOp(opID string, reqProps map[string]*openapi3.SchemaRef) *openapi3.PathItem {
	op := &openapi3.Operation{
		OperationID: opID,
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
					Type:       &openapi3.Types{"object"},
					Properties: reqProps,
				}),
			},
		},
		Responses: openapi3.NewResponses(),
	}
	return &openapi3.PathItem{Post: op}
}

// makeFS builds a Fullstack with the given pages and OpenAPI doc.
func makeFS(pages []stml.PageSpec, doc *openapi3.T) *yongol.Fullstack {
	return &yongol.Fullstack{
		SpecsDir:   "/tmp/test-project",
		STMLPages:  pages,
		OpenAPIDoc: doc,
	}
}

// stringProp creates a simple string schema ref.
func stringProp() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
}

// intProp creates a simple integer schema ref.
func intProp() *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}}
}

// arrayProp creates an array schema ref.
func arrayProp(itemType string) *openapi3.SchemaRef {
	return &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{itemType}}},
	}}
}

// paramRef creates a parameter ref.
func paramRef(name, in string) *openapi3.ParameterRef {
	return &openapi3.ParameterRef{Value: &openapi3.Parameter{Name: name, In: in}}
}

// hasDiag returns true if any diagnostic has the given rule prefix in its message.
func hasDiag(diags []diagnostic.Diagnostic, rulePrefix string) bool {
	for _, d := range diags {
		if len(d.Message) >= len(rulePrefix) && d.Message[:len(rulePrefix)] == rulePrefix {
			return true
		}
	}
	return false
}

// countDiag returns the number of diagnostics matching the given rule prefix.
func countDiag(diags []diagnostic.Diagnostic, rulePrefix string) int {
	n := 0
	for _, d := range diags {
		if len(d.Message) >= len(rulePrefix) && d.Message[:len(rulePrefix)] == rulePrefix {
			n++
		}
	}
	return n
}

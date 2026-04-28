//ff:func feature=validate type=test-helper control=sequence topic=hurl-statemachine
//ff:what newWorkflowDoc — 테스트용 workflow endpoint 들이 선언된 최소 OpenAPI 문서

package hurl_statemachine

import "github.com/getkin/kin-openapi/openapi3"

// newWorkflowDoc returns a minimal OpenAPI doc with the zenflow-like
// workflow endpoints used across the tests in this file.
func newWorkflowDoc() *openapi3.T {
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	post := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "CreateWorkflow", Responses: openapi3.NewResponses()}}
	doc.Paths.Set("/workflows", post)
	activate := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "ActivateWorkflow", Responses: openapi3.NewResponses()}}
	doc.Paths.Set("/workflows/{id}/activate", activate)
	execute := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "ExecuteWorkflow", Responses: openapi3.NewResponses()}}
	doc.Paths.Set("/workflows/{id}/execute", execute)
	return doc
}

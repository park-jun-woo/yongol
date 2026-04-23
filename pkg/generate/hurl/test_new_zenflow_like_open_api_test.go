//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what newZenflowLikeOpenAPI — ZenFlow Phase003/004 smoke 테스트용 *openapi3.T 조립

package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// newZenflowLikeOpenAPI returns an *openapi3.T modelling the ZenFlow
// subset needed by Phase003/004 smoke tests: auth (Register+Login),
// workflow CRUD (Create/List/Get), and the state-machine operations
// (AddAction / ActivateWorkflow / PauseWorkflow / ArchiveWorkflow /
// ExecuteWorkflow). Enough for the smoke walker to produce a full
// ordered scenario end-to-end.
func newZenflowLikeOpenAPI() *openapi3.T {
	workflowIDParam := &openapi3.ParameterRef{Value: &openapi3.Parameter{
		Name: "id",
		In:   "path",
	}}
	createWorkflow := &openapi3.Operation{
		OperationID: "CreateWorkflow",
		RequestBody: &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
			WithContent(openapi3.NewContentWithJSONSchema(&openapi3.Schema{
				Type: &openapi3.Types{"object"},
				Properties: openapi3.Schemas{
					"name": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				},
			}))},
		Responses: newCreatedWithIDResponses(),
	}
	listWorkflows := &openapi3.Operation{OperationID: "ListWorkflows", Responses: newOKResponses()}
	getWorkflow := &openapi3.Operation{
		OperationID: "GetWorkflow",
		Parameters:  openapi3.Parameters{workflowIDParam},
		Responses:   newOKResponses(),
	}
	addAction := &openapi3.Operation{
		OperationID: "AddAction",
		Parameters:  openapi3.Parameters{workflowIDParam},
		Responses:   newCreatedResponses(),
	}
	activate := &openapi3.Operation{OperationID: "ActivateWorkflow", Parameters: openapi3.Parameters{workflowIDParam}, Responses: newOKResponses()}
	pause := &openapi3.Operation{OperationID: "PauseWorkflow", Parameters: openapi3.Parameters{workflowIDParam}, Responses: newOKResponses()}
	archive := &openapi3.Operation{OperationID: "ArchiveWorkflow", Parameters: openapi3.Parameters{workflowIDParam}, Responses: newOKResponses()}
	execute := &openapi3.Operation{OperationID: "ExecuteWorkflow", Parameters: openapi3.Parameters{workflowIDParam}, Responses: newOKResponses()}

	doc := newAuthOnlyOpenAPI()
	authPaths := doc.Paths
	paths := openapi3.NewPaths(
		openapi3.WithPath("/auth/register", authPaths.Value("/auth/register")),
		openapi3.WithPath("/auth/login", authPaths.Value("/auth/login")),
		openapi3.WithPath("/workflows", &openapi3.PathItem{Post: createWorkflow, Get: listWorkflows}),
		openapi3.WithPath("/workflows/{id}", &openapi3.PathItem{Get: getWorkflow}),
		openapi3.WithPath("/workflows/{id}/actions", &openapi3.PathItem{Post: addAction}),
		openapi3.WithPath("/workflows/{id}/activate", &openapi3.PathItem{Post: activate}),
		openapi3.WithPath("/workflows/{id}/pause", &openapi3.PathItem{Post: pause}),
		openapi3.WithPath("/workflows/{id}/archive", &openapi3.PathItem{Post: archive}),
		openapi3.WithPath("/workflows/{id}/execute", &openapi3.PathItem{Post: execute}),
	)
	doc.Paths = paths
	return doc
}

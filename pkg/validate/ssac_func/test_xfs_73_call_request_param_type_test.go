//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what XFS-73 test — @call input request.* OpenAPI param type ↔ Func Request field type 검증

package ssac_func

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXfs73_UUIDParamVsStringFunc verifies that XFS-73 emits an ERROR when a
// path param with format:uuid (openapi_types.UUID) is passed to a Func Request
// field typed as string.
func TestXfs73_UUIDParamVsStringFunc(t *testing.T) {
	param := &openapi3.ParameterRef{Value: &openapi3.Parameter{
		Name: "id",
		In:   "path",
		Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type:   &openapi3.Types{"string"},
			Format: "uuid",
		}},
	}}
	op := &openapi3.Operation{
		OperationID: "HandleRequest",
		Parameters:  openapi3.Parameters{param},
	}
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/requests/{id}", &openapi3.PathItem{Post: op}),
	)}

	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "HandleRequest",
			FileName: "service/handle_request.ssac",
			Sequences: []parsessac.Sequence{{
				Type:  "call",
				Model: "report.GenerateReport",
				Line:  10,
				Inputs: map[string]string{
					"WorkflowID": "request.id",
				},
			}},
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{{
			Package: "report",
			Name:    "generateReport",
			RequestFields: []funcspec.Field{
				{Name: "WorkflowID", Type: "string"},
			},
			ReturnTypes: []string{"GenerateReportResponse", "error"},
		}},
	}
	fs.SetGround(ground.Build(fs))

	diags := xfs73CallRequestParamType(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XFS-73]") {
		t.Errorf("expected [XFS-73] in message, got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "openapi_types.UUID") {
		t.Errorf("expected openapi_types.UUID in message, got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "string") {
		t.Errorf("expected string in message, got %q", diags[0].Message)
	}
}

// TestXfs73_StringParamToStringFunc verifies that XFS-73 does NOT emit an
// error when a string param is passed to a string Func Request field.
func TestXfs73_StringParamToStringFunc(t *testing.T) {
	param := &openapi3.ParameterRef{Value: &openapi3.Parameter{
		Name: "slug",
		In:   "path",
		Schema: &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"string"},
		}},
	}}
	op := &openapi3.Operation{
		OperationID: "GetBySlug",
		Parameters:  openapi3.Parameters{param},
	}
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items/{slug}", &openapi3.PathItem{Get: op}),
	)}

	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "GetBySlug",
			FileName: "service/get_by_slug.ssac",
			Sequences: []parsessac.Sequence{{
				Type:  "call",
				Model: "item.FindBySlug",
				Line:  5,
				Inputs: map[string]string{
					"Slug": "request.slug",
				},
			}},
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{{
			Package: "item",
			Name:    "findBySlug",
			RequestFields: []funcspec.Field{
				{Name: "Slug", Type: "string"},
			},
			ReturnTypes: []string{"FindBySlugResponse", "error"},
		}},
	}
	fs.SetGround(ground.Build(fs))

	diags := xfs73CallRequestParamType(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diagnostics for string→string, got %d (%+v)", len(diags), diags)
	}
}

// TestXfs73_BodyFieldSkipped verifies that XFS-73 skips request.* inputs that
// refer to body fields (not registered as OpenAPI.paramType).
func TestXfs73_BodyFieldSkipped(t *testing.T) {
	// No OpenAPI params — "title" is a body field, not a path/query param.
	op := &openapi3.Operation{
		OperationID: "CreateItem",
	}
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items", &openapi3.PathItem{Post: op}),
	)}

	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "CreateItem",
			FileName: "service/create_item.ssac",
			Sequences: []parsessac.Sequence{{
				Type:  "call",
				Model: "item.ProcessItem",
				Line:  3,
				Inputs: map[string]string{
					"Title": "request.title",
				},
			}},
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{{
			Package: "item",
			Name:    "processItem",
			RequestFields: []funcspec.Field{
				{Name: "Title", Type: "int64"},
			},
			ReturnTypes: []string{"ProcessItemResponse", "error"},
		}},
	}
	fs.SetGround(ground.Build(fs))

	diags := xfs73CallRequestParamType(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diagnostics for body field, got %d (%+v)", len(diags), diags)
	}
}

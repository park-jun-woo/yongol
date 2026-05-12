//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestXfs73_UUIDParamVsStringFunc — UUID param vs string Func 필드 타입 불일치 XFS-73 발생 검증

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

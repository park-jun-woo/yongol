//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestXfs73_BodyFieldSkipped — body field 는 XFS-73 검사 스킵 검증

package ssac_func

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXfs73_BodyFieldSkipped verifies that XFS-73 skips request.* inputs that
// refer to body fields (not registered as OpenAPI.paramType).
func TestXfs73_BodyFieldSkipped(t *testing.T) {
	// No OpenAPI params -- "title" is a body field, not a path/query param.
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

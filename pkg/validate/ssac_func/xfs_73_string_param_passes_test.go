//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestXfs73_StringParamToStringFunc — string param → string Func 필드는 XFS-73 미발생 검증

package ssac_func

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
		t.Errorf("expected 0 diagnostics for string->string, got %d (%+v)", len(diags), diags)
	}
}

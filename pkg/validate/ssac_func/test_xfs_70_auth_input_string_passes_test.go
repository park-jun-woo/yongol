//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestXfs70_AuthInputStringPasses — string 타입 @auth input 은 XFS-70 미발생 검증

package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXfs70_AuthInputStringPasses verifies that @auth with string-typed
// input (e.g. request.id path param) does not trigger XFS-70.
func TestXfs70_AuthInputStringPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "GetWorkflow",
			FileName: "service/get_workflow.ssac",
			Sequences: []parsessac.Sequence{{
				Type:     "auth",
				Action:   "GetWorkflow",
				Resource: "workflow",
				Line:     3,
				Inputs: map[string]string{
					"ResourceID": "request.id",
				},
			}},
		}},
	}
	fs.SetGround(ground.Build(fs))
	diags := xfs70AuthInputType(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diagnostics for string input, got %d (%+v)", len(diags), diags)
	}
}

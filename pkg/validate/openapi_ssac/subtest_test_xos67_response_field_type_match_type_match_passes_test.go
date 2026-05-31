//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXos67ResponseFieldTypeMatchTypeMatchPasses — type match passes 서브테스트
package openapi_ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func subtestTestXos67ResponseFieldTypeMatchTypeMatchPasses(t *testing.T) {

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name: "getUser",
				Sequences: []ssac.Sequence{
					{Type: "get", Result: &ssac.Result{Var: "user", Type: "User"}},
					{Type: "response", Fields: map[string]string{"name": "user.Name"}},
				},
			},
		},
	}
	g := &rule.Ground{
		Types: map[string]string{
			"OpenAPI.response.getUser.name": "string",
			"SSaC.var.getUser.user":         "User",
			"Struct.User.Name":              "string",
		},
	}
	fs.SetGround(g)
	diags := xos67ResponseFieldType(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}

}

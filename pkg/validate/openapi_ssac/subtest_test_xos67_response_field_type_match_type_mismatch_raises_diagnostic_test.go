//ff:func feature=validate type=test-helper control=sequence
//ff:what subtestTestXos67ResponseFieldTypeMatchTypeMismatchRaisesDiagnostic — type mismatch raises diagnostic 서브테스트
package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func subtestTestXos67ResponseFieldTypeMatchTypeMismatchRaisesDiagnostic(t *testing.T) {

	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{
			{
				Name:     "getUser",
				FileName: "user.ssac",
				Sequences: []ssac.Sequence{
					{Type: "get", Result: &ssac.Result{Var: "user", Type: "User"}},
					{Type: "response", Fields: map[string]string{"name": "user.ID"}},
				},
			},
		},
	}
	g := &rule.Ground{
		Types: map[string]string{
			"OpenAPI.response.getUser.name": "string",
			"SSaC.var.getUser.user":         "User",
			"Struct.User.ID":                "int64",
		},
	}
	fs.SetGround(g)
	diags := xos67ResponseFieldType(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "XOS-67") {
		t.Errorf("Message missing XOS-67: %s", diags[0].Message)
	}

}

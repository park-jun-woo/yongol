//ff:func feature=validate type=test control=sequence topic=authz-check
//ff:what XAS-60 positive — @auth input key 가 CheckRequest 필드 집합에 없으면 ERROR

package ssac_authz

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXas60AuthInputFieldFires(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:     "DeleteProject",
			FileName: "project.ssac",
			Sequences: []ssac.Sequence{{
				Type:     "auth",
				Line:     42,
				Action:   "delete",
				Resource: "project",
				Inputs: map[string]string{
					// "subject_id" and "resource_id" are legitimate CheckRequest
					// fields; "bogus" is not.
					"subject_id":  "user.ID",
					"resource_id": "project.ID",
					"bogus":       "project.X",
				},
			}},
		}},
	}
	fs.SetGround(&rule.Ground{
		Lookup: map[string]rule.StringSet{
			"Authz.checkRequest": {
				"subject_id":  true,
				"resource_id": true,
			},
		},
	})

	diags := xas60AuthInputField(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XAS-60]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "bogus") {
		t.Errorf("expected unknown key 'bogus' in message, got %q", diags[0].Message)
	}
}

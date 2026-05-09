//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what XFS-44 test — resolveInputType currentUser 필드 해석 + uuid claim vs string Func 필드 타입 불일치 진단

package ssac_func

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestResolveInputType_CurrentUserField verifies that resolveInputType returns
// the registered claim type for currentUser.<Field> expressions.
func TestResolveInputType_CurrentUserField(t *testing.T) {
	g := &rule.Ground{
		Types: map[string]string{
			"Manifest.claim.OrgID":  "pgtype.UUID",
			"Manifest.claim.UserID": "int64",
		},
	}
	tests := []struct {
		value string
		want  string
	}{
		{"currentUser.OrgID", "pgtype.UUID"},
		{"currentUser.UserID", "int64"},
		{"currentUser.Unknown", ""},
		{"org.Name", ""},
	}
	for _, tt := range tests {
		got := resolveInputType(g, "anyFunc", tt.value)
		if got != tt.want {
			t.Errorf("resolveInputType(g, %q, %q) = %q, want %q", "anyFunc", tt.value, got, tt.want)
		}
	}
}

// TestXfs44_ClaimUUIDvsStringFunc verifies that XFS-44 emits an ERROR when a
// uuid claim field (pgtype.UUID) is passed to a Func Request field typed as
// string.
func TestXfs44_ClaimUUIDvsStringFunc(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Auth: &manifest.Auth{
					Type: "jwt",
					Claims: map[string]manifest.ClaimDef{
						"OrgID": {Key: "org_id", GoType: "uuid"},
					},
				},
			},
		},
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "HandleRequest",
			FileName: "service/handle_request.ssac",
			Sequences: []parsessac.Sequence{{
				Type:  "call",
				Model: "billing.CheckCredits",
				Line:  10,
				Inputs: map[string]string{
					"OrgID": "currentUser.OrgID",
				},
			}},
		}},
		ProjectFuncSpecs: []funcspec.FuncSpec{{
			Package: "billing",
			Name:    "checkCredits",
			RequestFields: []funcspec.Field{
				{Name: "OrgID", Type: "string"},
			},
			ReturnTypes: []string{"CheckCreditsResponse", "error"},
		}},
	}
	fs.SetGround(ground.Build(fs))
	diags := xfs44CallInputType(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XFS-44]") {
		t.Errorf("expected [XFS-44] in message, got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "pgtype.UUID") {
		t.Errorf("expected pgtype.UUID in message, got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "string") {
		t.Errorf("expected string in message, got %q", diags[0].Message)
	}
}

//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos66UsedFieldsRequired — no constraints/required/not required 검증

package openapi_ssac

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos66UsedFieldsRequired_Unit(t *testing.T) {
	t.Run("no constraints returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{{Name: "getUser"}},
		}
		diags := xos66UsedFieldsRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("required field passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "createUser",
					Sequences: []ssac.Sequence{
						{Type: "post", Args: []ssac.Arg{{Source: "request", Field: "Email"}}},
					},
				},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"Email": {Required: true}},
			},
		}
		diags := xos66UsedFieldsRequired(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("optional field raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "createUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "post", Args: []ssac.Arg{{Source: "request", Field: "Email"}}},
					},
				},
			},
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {"Email": {Required: false}},
			},
		}
		diags := xos66UsedFieldsRequired(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XOS-66") {
			t.Errorf("Message missing XOS-66: %s", diags[0].Message)
		}
	})
}

//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos70_ConstraintMissing — 필드/함수 constraint 미존재 시 스킵 검증

package openapi_ssac

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos70_ConstraintMissing(t *testing.T) {
	t.Run("FieldNotInConstraints_Skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{
					{Type: "response", Fields: map[string]string{"count": "0"}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"getUser": {},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("NoConstraintsForFunc_Skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{
					{Type: "response", Fields: map[string]string{"count": "0"}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}

//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos70_SkipCases — non-response/non-integer literal/required/non-integer type 스킵 검증

package openapi_ssac

import (
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos70_SkipCases(t *testing.T) {
	t.Run("NonResponseSeqSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{
					{Type: "get", Fields: map[string]string{"count": "0"}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"getUser": {"count": {Type: "integer", Format: "", Required: false}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("NonIntegerLiteralSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{
					{Type: "response", Fields: map[string]string{"name": `"hello"`}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"getUser": {"name": {Type: "string", Required: false}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("RequiredFieldSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{
					{Type: "response", Fields: map[string]string{"count": "0"}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"getUser": {"count": {Type: "integer", Format: "", Required: true}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("NonIntegerTypeSkipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{
					{Type: "response", Fields: map[string]string{"count": "0"}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"getUser": {"count": {Type: "string", Format: "", Required: false}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}

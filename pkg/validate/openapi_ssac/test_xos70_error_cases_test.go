//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos70_ErrorCases — optional integer + format 미지정/int32 시 진단 발생 검증

package openapi_ssac

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos70_ErrorCases(t *testing.T) {
	t.Run("OptionalIntegerWithoutInt64_Error", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "getUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Line: 10, Fields: map[string]string{"count": "0"}},
					},
				},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"getUser": {"count": {Type: "integer", Format: "", Required: false}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		d := diags[0]
		if !strings.Contains(d.Message, "XOS-70") {
			t.Errorf("message should contain XOS-70: %s", d.Message)
		}
		if !strings.Contains(d.Message, "count") {
			t.Errorf("message should contain field name: %s", d.Message)
		}
		if d.File != "user.ssac" {
			t.Errorf("file = %q, want user.ssac", d.File)
		}
		if d.Line != 10 {
			t.Errorf("line = %d, want 10", d.Line)
		}
		if d.OperationID != "getUser" {
			t.Errorf("operationID = %q, want getUser", d.OperationID)
		}
	})

	t.Run("OptionalIntegerInt32_Error", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "getItem",
					FileName: "item.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Line: 5, Fields: map[string]string{"qty": "1"}},
					},
				},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"getItem": {"qty": {Type: "integer", Format: "int32", Required: false}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XOS-70") {
			t.Errorf("message should contain XOS-70: %s", diags[0].Message)
		}
	})
}

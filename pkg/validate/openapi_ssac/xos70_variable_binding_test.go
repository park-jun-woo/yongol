//ff:func feature=validate type=test control=sequence topic=openapi-ssac
//ff:what TestXos70_VariableBinding — 변수 바인딩 정수 응답(비-DDL COUNT 등)의 format: int64 강제 검증

package openapi_ssac

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos70_VariableBinding(t *testing.T) {
	t.Run("OptionalVariableWithoutInt64_Error", func(t *testing.T) {
		// non-DDL COUNT result bound to an optional integer field; codegen emits
		// &count (*int64) but formatless oapi-codegen *int → mismatch.
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "listUsers", FileName: "user.ssac", Sequences: []ssac.Sequence{
					{Type: "response", Line: 7, Fields: map[string]string{"total": "count"}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"listUsers": {"total": {Type: "integer", Format: "", Required: false}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "variable binding") {
			t.Errorf("message should mention variable binding: %s", diags[0].Message)
		}
	})

	t.Run("RequiredVariableWithoutInt64_Error", func(t *testing.T) {
		// required integer field bound to a variable → int64 directly, formatless int → mismatch.
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "listUsers", FileName: "user.ssac", Sequences: []ssac.Sequence{
					{Type: "response", Line: 7, Fields: map[string]string{"total": "count"}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"listUsers": {"total": {Type: "integer", Format: "", Required: true}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("VariableWithInt64_Passes", func(t *testing.T) {
		// DDL-backed / properly-typed integer field already carries format: int64.
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "listUsers", FileName: "user.ssac", Sequences: []ssac.Sequence{
					{Type: "response", Line: 7, Fields: map[string]string{"total": "count"}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"listUsers": {"total": {Type: "integer", Format: "int64", Required: false}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("StringLiteralToIntegerField_SkippedForXos67", func(t *testing.T) {
		// string literal mapped to an integer field is a type mismatch (XOS-67),
		// not a width issue — XOS-70 must not double-report it.
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "listUsers", FileName: "user.ssac", Sequences: []ssac.Sequence{
					{Type: "response", Line: 7, Fields: map[string]string{"total": `"oops"`}},
				}},
			},
			ResponseConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"listUsers": {"total": {Type: "integer", Format: "", Required: false}},
			},
		}
		diags := xos70ResponseLiteralIntFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}

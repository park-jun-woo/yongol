//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what xoo71PasswordNoMinLength — empty constraints + password 필드 minLength 유무 검증
package openapi

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoo71PasswordNoMinLength(t *testing.T) {
	minLen := 8

	t.Run("empty request constraints returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xoo71PasswordNoMinLength(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("non-password field skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {
					"username": {Type: "string"},
				},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xoo71PasswordNoMinLength(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("password with minLength passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {
					"password": {Type: "string", MinLength: &minLen},
				},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xoo71PasswordNoMinLength(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("password without minLength raises warning", func(t *testing.T) {
		fs := &yongol.Fullstack{
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {
					"password": {Type: "string"},
				},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xoo71PasswordNoMinLength(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XOO-71") {
			t.Errorf("Message missing XOO-71: %s", diags[0].Message)
		}
	})
}

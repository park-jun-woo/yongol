//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what xoo72EmailNoFormat — empty constraints + 비email skip + format 유무 검증

package openapi

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXoo72EmailNoFormat_Unit(t *testing.T) {
	t.Run("empty constraints returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xoo72EmailNoFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("non-email field skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"createUser": {
					"username": {Type: "string"},
				},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xoo72EmailNoFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("email with format passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"register": {
					"email": {Type: "string", Format: "email"},
				},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xoo72EmailNoFormat(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("email without format raises warning", func(t *testing.T) {
		fs := &yongol.Fullstack{
			RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
				"register": {
					"email": {Type: "string"},
				},
			},
			OpenAPILines: &oapiparser.LineIndex{RequestFields: map[string]map[string]int{}},
		}
		diags := xoo72EmailNoFormat(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XOO-72") {
			t.Errorf("Message missing XOO-72: %s", diags[0].Message)
		}
	})
}

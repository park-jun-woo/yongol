//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos67ResponseFieldType match — type match/mismatch/unresolvable 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos67ResponseFieldType_Match(t *testing.T) {
	t.Run("type match passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "getUser",
					Sequences: []ssac.Sequence{
						{Type: "get", Result: &ssac.Result{Var: "user", Type: "User"}},
						{Type: "response", Fields: map[string]string{"name": "user.Name"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				"OpenAPI.response.getUser.name": "string",
				"SSaC.var.getUser.user":         "User",
				"Struct.User.Name":              "string",
			},
		}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("type mismatch raises diagnostic", func(t *testing.T) {
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
	})

	t.Run("unresolvable expected type skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "getUser",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{"name": `"hello"`}},
					},
				},
			},
		}
		g := &rule.Ground{Types: map[string]string{}}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("timestamptz bound to string+date-time passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "approveItem",
					Sequences: []ssac.Sequence{
						{Type: "put", Result: &ssac.Result{Var: "updated", Type: "Item"}},
						{Type: "response", Fields: map[string]string{"approved_at": "updated.ApprovedAt"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				"OpenAPI.response.approveItem.approved_at":        "string",
				"OpenAPI.response.approveItem.approved_at.format": "date-time",
				"SSaC.var.approveItem.updated":                    "Item",
				"Struct.Item.ApprovedAt":                          "time.Time",
			},
		}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("string without date-time vs time.Time still errors", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "approveItem",
					FileName: "item.ssac",
					Sequences: []ssac.Sequence{
						{Type: "put", Result: &ssac.Result{Var: "updated", Type: "Item"}},
						{Type: "response", Fields: map[string]string{"approved_at": "updated.ApprovedAt"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				"OpenAPI.response.approveItem.approved_at": "string",
				"SSaC.var.approveItem.updated":             "Item",
				"Struct.Item.ApprovedAt":                   "time.Time",
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
	})

	t.Run("string literal bound to string field passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "ping",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{"status": `"ok"`}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				"OpenAPI.response.ping.status": "string",
			},
		}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("unresolvable actual type skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "getUser",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{"name": "unknown.Field"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				"OpenAPI.response.getUser.name": "string",
			},
		}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
}

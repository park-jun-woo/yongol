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

	// Format-aware (Phase018): a { type: string, format: date-time } response
	// field is registered as expected=time.Time (registerOpenAPIResponseProps),
	// so a DDL TIMESTAMPTZ column (time.Time) matches directly. The old
	// expected="string" + ".format" marker allowance has been removed.
	t.Run("timestamptz bound to date-time field passes", func(t *testing.T) {
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
				"OpenAPI.response.approveItem.approved_at": "time.Time",
				"SSaC.var.approveItem.updated":             "Item",
				"Struct.Item.ApprovedAt":                   "time.Time",
			},
		}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("string literal bound to date-time field errors (false negative closed)", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "approveItem",
					FileName: "item.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{"approved_at": `"2026-01-01T00:00:00Z"`}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				// expected is time.Time (format: date-time); a string literal
				// (actual="string") must now ERROR — previously slipped through.
				"OpenAPI.response.approveItem.approved_at": "time.Time",
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

	// --- Phase018 format:uuid response field cases ---

	t.Run("uuid field bound to DB UUID column passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "getWorkflow",
					Sequences: []ssac.Sequence{
						{Type: "get", Result: &ssac.Result{Var: "wf", Type: "Workflow"}},
						{Type: "response", Fields: map[string]string{"id": "wf.ID"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				// expected: format:uuid → openapi_types.UUID
				"OpenAPI.response.getWorkflow.id": "openapi_types.UUID",
				"SSaC.var.getWorkflow.wf":         "Workflow",
				// actual: DB UUID column resolved via api-surface type
				"DDL.apifield.Workflow.ID": "openapi_types.UUID",
				// coarse GoTypeOf projection (collapses UUID→string) — must be
				// overridden by the apifield key above.
				"Struct.Workflow.ID": "string",
			},
		}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("uuid field bound to func openapi_types.UUID passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "cancelMatch",
					Sequences: []ssac.Sequence{
						{Type: "put", Result: &ssac.Result{Var: "result", Type: "CancelMatchResponse"}},
						{Type: "response", Fields: map[string]string{"match_id": "result.ID"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				"OpenAPI.response.cancelMatch.match_id": "openapi_types.UUID",
				"SSaC.var.cancelMatch.result":           "CancelMatchResponse",
				// func struct fields register raw declared type (no apifield key).
				"Struct.CancelMatchResponse.ID": "openapi_types.UUID",
			},
		}
		fs.SetGround(g)
		diags := xos67ResponseFieldType(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("uuid field bound to func pgtype.UUID errors with expected openapi_types.UUID", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "cancelMatch",
					FileName: "cancel_match.ssac",
					Sequences: []ssac.Sequence{
						{Type: "put", Result: &ssac.Result{Var: "result", Type: "CancelMatchResponse"}},
						{Type: "response", Fields: map[string]string{"match_id": "result.ID"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				"OpenAPI.response.cancelMatch.match_id": "openapi_types.UUID",
				"SSaC.var.cancelMatch.result":           "CancelMatchResponse",
				// func field wrongly declared as DB/sqlc type pgtype.UUID.
				"Struct.CancelMatchResponse.ID": "pgtype.UUID",
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
		if !strings.Contains(diags[0].Message, "openapi_types.UUID") {
			t.Errorf("Message should mention expected openapi_types.UUID: %s", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "pgtype.UUID") {
			t.Errorf("Message should mention actual pgtype.UUID: %s", diags[0].Message)
		}
	})

	t.Run("uuid field bound to string literal errors (false negative closed)", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "cancelMatch",
					FileName: "cancel_match.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{"match_id": `"abc"`}},
					},
				},
			},
		}
		g := &rule.Ground{
			Types: map[string]string{
				"OpenAPI.response.cancelMatch.match_id": "openapi_types.UUID",
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
}

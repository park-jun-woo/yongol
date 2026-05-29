//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what xos17ResponseFields — nil ground/no keys/매칭/누락 검증

package openapi_ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXos17ResponseFields_Unit(t *testing.T) {
	t.Run("nil ground returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := xos17ResponseFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no explicit response keys skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{Name: "getUser", Sequences: []ssac.Sequence{{Type: "get"}}},
			},
		}
		g := &rule.Ground{Schemas: map[string][]string{}}
		fs.SetGround(g)
		diags := xos17ResponseFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("empty OpenAPI response schema skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "getUser",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{"id": "user.ID"}},
					},
				},
			},
		}
		g := &rule.Ground{Schemas: map[string][]string{}}
		fs.SetGround(g)
		diags := xos17ResponseFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})

	t.Run("matching fields passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name: "getUser",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{"id": "user.ID", "email": "user.Email"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{
				"OpenAPI.response.getUser": {"id", "email"},
			},
		}
		fs.SetGround(g)
		diags := xos17ResponseFields(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing field raises diagnostic", func(t *testing.T) {
		fs := &yongol.Fullstack{
			ServiceFuncs: []ssac.ServiceFunc{
				{
					Name:     "getUser",
					FileName: "user.ssac",
					Sequences: []ssac.Sequence{
						{Type: "response", Fields: map[string]string{"phone": "user.Phone"}},
					},
				},
			},
		}
		g := &rule.Ground{
			Schemas: map[string][]string{
				"OpenAPI.response.getUser": {"id", "email"},
			},
		}
		fs.SetGround(g)
		diags := xos17ResponseFields(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "XOS-17") {
			t.Errorf("Message missing XOS-17: %s", diags[0].Message)
		}
	})
}

func TestXos17ResponseFields(t *testing.T) {
	_ = t
	_ = &yongol.Fullstack{}
}

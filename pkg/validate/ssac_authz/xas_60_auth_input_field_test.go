//ff:func feature=validate type=test control=sequence topic=authz-check
//ff:what XAS-60 — invalid @auth field fires, valid passes, nil ground/empty/non-auth skip

package ssac_authz

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXas60AuthInputField(t *testing.T) {
	authSeq := func(inputs map[string]string) parsessac.Sequence {
		return parsessac.Sequence{Type: "auth", Action: "Delete", Resource: "project", Inputs: inputs, Line: 5}
	}
	t.Run("fires_on_unknown", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{Name: "Del", FileName: "s/del.ssac", Sequences: []parsessac.Sequence{authSeq(map[string]string{"BadField": "p.ID"})}}}}
		fs.SetGround(xas60Ground("ResourceID", "Action"))
		diags := xas60AuthInputField(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d", len(diags))
		}
		if !strings.Contains(diags[0].Message, "[XAS-60]") {
			t.Errorf("prefix: %q", diags[0].Message)
		}
	})
	t.Run("passes_on_valid", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{Name: "Del", FileName: "s/del.ssac", Sequences: []parsessac.Sequence{authSeq(map[string]string{"ResourceID": "p.ID"})}}}}
		fs.SetGround(xas60Ground("ResourceID", "Action"))
		diags := xas60AuthInputField(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("skips_nil_ground", func(t *testing.T) {
		diags := xas60AuthInputField(&yongol.Fullstack{})
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("skips_empty_fields", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{Name: "Del", FileName: "s/del.ssac", Sequences: []parsessac.Sequence{authSeq(map[string]string{"X": "y"})}}}}
		fs.SetGround(&rule.Ground{Lookup: map[string]rule.StringSet{}, Types: map[string]string{}, Pairs: map[string]rule.StringSet{}, Config: map[string]bool{}, Vars: rule.StringSet{}, Flags: rule.StringSet{}})
		diags := xas60AuthInputField(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
	t.Run("skips_non_auth", func(t *testing.T) {
		fs := &yongol.Fullstack{ServiceFuncs: []parsessac.ServiceFunc{{Name: "Get", FileName: "s/g.ssac", Sequences: []parsessac.Sequence{{Type: "get", Inputs: map[string]string{"Bad": "x"}, Line: 5}}}}}
		fs.SetGround(xas60Ground("ResourceID"))
		diags := xas60AuthInputField(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d", len(diags))
		}
	})
}

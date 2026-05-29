//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what findOwnershipMapping 단위 테스트 (resource 일치하는 첫 매핑 반환)

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

func TestFindOwnershipMapping(t *testing.T) {
	owns := []rego.OwnershipMapping{
		{Resource: "workflow", Table: "workflows"},
		{Resource: "project", Table: "projects"},
		{Resource: "workflow", Table: "workflows_v2"}, // duplicate resource
	}
	t.Run("returns first match", func(t *testing.T) {
		m := findOwnershipMapping(owns, "workflow")
		if m == nil || m.Table != "workflows" {
			t.Fatalf("expected workflows mapping, got %+v", m)
		}
	})
	t.Run("returns project match", func(t *testing.T) {
		m := findOwnershipMapping(owns, "project")
		if m == nil || m.Table != "projects" {
			t.Fatalf("expected projects mapping, got %+v", m)
		}
	})
	t.Run("no match → nil", func(t *testing.T) {
		if m := findOwnershipMapping(owns, "ghost"); m != nil {
			t.Errorf("expected nil, got %+v", m)
		}
	})
}

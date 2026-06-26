//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildSyntheticFuncSpec — 내부 타입에 대한 합성 FuncSpec 생성 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestBuildSyntheticFuncSpec(t *testing.T) {
	pkgTypes := map[string]map[string][]funcspec.Field{
		"chat": {
			"ChatMessage": {
				{Name: "TurnID", Type: "string", JSONName: "turn_id"},
				{Name: "Content", Type: "string", JSONName: "content"},
			},
		},
	}

	t.Run("Found", func(t *testing.T) {
		spec := buildSyntheticFuncSpec("ChatMessage", "chat", pkgTypes)
		if spec == nil {
			t.Fatal("expected non-nil spec")
		}
		if spec.Package != "chat" || spec.Name != "ChatMessage" {
			t.Errorf("got Package=%q Name=%q", spec.Package, spec.Name)
		}
		if len(spec.ResponseFields) != 2 || spec.ResponseFields[0].Name != "TurnID" {
			t.Errorf("ResponseFields = %+v", spec.ResponseFields)
		}
	})

	t.Run("NilPkgTypes", func(t *testing.T) {
		if spec := buildSyntheticFuncSpec("ChatMessage", "chat", nil); spec != nil {
			t.Errorf("expected nil when funcPackageTypes is nil")
		}
	})

	t.Run("MissingPackage", func(t *testing.T) {
		other := map[string]map[string][]funcspec.Field{"auth": {}}
		if spec := buildSyntheticFuncSpec("ChatMessage", "chat", other); spec != nil {
			t.Errorf("expected nil when package not found")
		}
	})

	t.Run("MissingType", func(t *testing.T) {
		other := map[string]map[string][]funcspec.Field{
			"chat": {"OtherType": {{Name: "X", Type: "string"}}},
		}
		if spec := buildSyntheticFuncSpec("ChatMessage", "chat", other); spec != nil {
			t.Errorf("expected nil when type not found in package")
		}
	})
}

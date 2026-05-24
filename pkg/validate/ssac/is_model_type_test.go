//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what isModelType — Model/Struct 판정 (nil ground/DDL model/struct/not found) 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestIsModelType(t *testing.T) {
	t.Run("nil ground returns false", func(t *testing.T) {
		if isModelType(nil, "User") {
			t.Error("expected false")
		}
	})

	t.Run("DDL model found", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"SymbolTable.model": {"User": true},
			},
		}
		if !isModelType(g, "User") {
			t.Error("expected true for DDL model")
		}
	})

	t.Run("struct found in Types", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{},
			Types:  map[string]string{"Struct.MyResponse.Name": "string"},
		}
		if !isModelType(g, "MyResponse") {
			t.Error("expected true for struct in Types")
		}
	})

	t.Run("not found", func(t *testing.T) {
		g := &rule.Ground{
			Lookup: map[string]rule.StringSet{
				"SymbolTable.model": {"Order": true},
			},
			Types: map[string]string{},
		}
		if isModelType(g, "User") {
			t.Error("expected false")
		}
	})
}

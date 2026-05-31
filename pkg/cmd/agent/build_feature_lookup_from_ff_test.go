//ff:func feature=agent type=test control=sequence
//ff:what TestBuildFeatureLookupFromFF — nil/빈/다수 feature 의 op→Feature 맵 구성 검증
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildFeatureLookupFromFF(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		if got := buildFeatureLookupFromFF(nil); got != nil {
			t.Errorf("nil input should return nil, got %v", got)
		}
	})

	t.Run("Populated", func(t *testing.T) {
		ff := &features.FeaturesFile{
			Features: []features.Feature{
				{Op: "CreateThing", Desc: "c"},
				{Op: "DeleteThing", Desc: "d"},
			},
		}
		m := buildFeatureLookupFromFF(ff)
		if len(m) != 2 {
			t.Fatalf("expected 2 entries, got %d", len(m))
		}
		if m["CreateThing"].Desc != "c" {
			t.Errorf("CreateThing desc: got %q, want c", m["CreateThing"].Desc)
		}
		if m["DeleteThing"].Desc != "d" {
			t.Errorf("DeleteThing desc: got %q, want d", m["DeleteThing"].Desc)
		}
		if _, ok := m["Missing"]; ok {
			t.Error("unexpected entry for Missing")
		}
	})

	t.Run("Empty", func(t *testing.T) {
		m := buildFeatureLookupFromFF(&features.FeaturesFile{})
		if len(m) != 0 {
			t.Errorf("expected empty map, got %v", m)
		}
	})
}

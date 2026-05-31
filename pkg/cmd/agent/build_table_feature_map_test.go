//ff:func feature=agent type=test control=sequence
//ff:what TestBuildTableFeatureMap — table 기준 feature 그룹화 및 빈 table skip 검증
package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildTableFeatureMap(t *testing.T) {
	ff := &features.FeaturesFile{
		Features: []features.Feature{
			{Op: "CreateWorkflow", Table: "workflows"},
			{Op: "GetWorkflow", Table: "workflows"},
			{Op: "ListProjects", Table: "projects"},
			{Op: "Healthz", Table: ""}, // skipped (no table)
		},
	}
	m := buildTableFeatureMap(ff)

	if len(m) != 2 {
		t.Fatalf("expected 2 table groups, got %d: %v", len(m), m)
	}
	if len(m["workflows"]) != 2 {
		t.Errorf("workflows: expected 2 features, got %d", len(m["workflows"]))
	}
	if len(m["projects"]) != 1 {
		t.Errorf("projects: expected 1 feature, got %d", len(m["projects"]))
	}
	if _, ok := m[""]; ok {
		t.Error("feature with empty table must be skipped")
	}

	t.Run("Empty", func(t *testing.T) {
		got := buildTableFeatureMap(&features.FeaturesFile{})
		if len(got) != 0 {
			t.Errorf("expected empty map, got %v", got)
		}
	})
}

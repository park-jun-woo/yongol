//ff:func feature=agent type=test control=sequence
//ff:what TestBuildDDLUserPrompt — 관계/상태/관련기능 유무에 따른 DDL prompt 구성 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildDDLUserPrompt(t *testing.T) {
	t.Run("FullTableDefWithFeatures", func(t *testing.T) {
		td := features.TableDef{
			BelongsTo: []string{"users", "orgs"},
			HasMany:   []string{"steps"},
			States:    []string{"draft", "active"},
		}
		feats := []features.Feature{
			{Op: "CreateWorkflow", Path: "/v1/workflows", Desc: "create"},
		}
		got := buildDDLUserPrompt("workflows", td, feats)
		for _, want := range []string{
			"Table: workflows",
			"belongs_to: users, orgs",
			"has_many: steps",
			"states: draft, active",
			"Related features:",
			"- CreateWorkflow /v1/workflows: create",
			"CREATE TABLE",
		} {
			if !strings.Contains(got, want) {
				t.Errorf("expected output to contain %q, got:\n%s", want, got)
			}
		}
	})

	t.Run("MinimalTableDefNoFeatures", func(t *testing.T) {
		got := buildDDLUserPrompt("plain", features.TableDef{}, nil)
		if !strings.Contains(got, "Table: plain") {
			t.Errorf("expected table line, got:\n%s", got)
		}
		for _, unwanted := range []string{"belongs_to:", "has_many:", "states:", "Related features:"} {
			if strings.Contains(got, unwanted) {
				t.Errorf("did not expect %q for minimal table def, got:\n%s", unwanted, got)
			}
		}
	})
}

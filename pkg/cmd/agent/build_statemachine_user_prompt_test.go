//ff:func feature=agent type=test control=sequence
//ff:what TestBuildStateMachineUserPrompt — 관련 feature 유무에 따른 state diagram 프롬프트 구성 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildStateMachineUserPrompt(t *testing.T) {
	states := []string{"draft", "active", "archived"}

	t.Run("WithFeatures", func(t *testing.T) {
		feats := []features.Feature{
			{Op: "ArchiveWorkflow", Path: "/v1/workflows/{id}/archive", Desc: "archive"},
		}
		got := buildStateMachineUserPrompt("workflows", states, feats)
		for _, want := range []string{
			"Table: workflows",
			"States: draft, active, archived",
			"Related features:",
			"- ArchiveWorkflow /v1/workflows/{id}/archive: archive",
			"Mermaid stateDiagram-v2",
		} {
			if !strings.Contains(got, want) {
				t.Errorf("expected %q, got:\n%s", want, got)
			}
		}
	})

	t.Run("WithoutFeatures", func(t *testing.T) {
		got := buildStateMachineUserPrompt("workflows", states, nil)
		if strings.Contains(got, "Related features:") {
			t.Errorf("did not expect features section, got:\n%s", got)
		}
		if !strings.Contains(got, "States: draft, active, archived") {
			t.Errorf("expected states line, got:\n%s", got)
		}
	})
}

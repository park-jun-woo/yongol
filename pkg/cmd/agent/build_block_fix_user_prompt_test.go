//ff:func feature=agent type=test control=sequence
//ff:what TestBuildBlockFixUserPrompt — desc 조회/layer 분기/진단 출력 등 prompt 구성 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildBlockFixUserPrompt(t *testing.T) {
	specsDir := t.TempDir()
	ff := &features.FeaturesFile{
		Features: []features.Feature{
			{Op: "CreateWorkflow", Desc: "create a workflow", Table: "workflows", Public: false},
		},
	}
	diags := []diagnostic.Diagnostic{
		{Message: "X-99: bad block", Advice: "fix it"},
	}

	t.Run("DescLookedUpAndOpenAPILayer", func(t *testing.T) {
		// Empty desc → resolved from the feature lookup; layerOpenAPI branch runs.
		got := buildBlockFixUserPrompt(specsDir, ff, "CreateWorkflow", "", "BLOCK_BODY", diags, layerOpenAPI)
		if !strings.Contains(got, "Feature: create a workflow") {
			t.Errorf("expected resolved feature desc, got:\n%s", got)
		}
		if !strings.Contains(got, "OperationId: CreateWorkflow") {
			t.Errorf("expected operationId line, got:\n%s", got)
		}
		if !strings.Contains(got, "Current block:\nBLOCK_BODY") {
			t.Errorf("expected current block, got:\n%s", got)
		}
		if !strings.Contains(got, "X-99: bad block") || !strings.Contains(got, "Advice: fix it") {
			t.Errorf("expected diagnostics rendered, got:\n%s", got)
		}
		if !strings.Contains(got, "Fix ONLY this block.") {
			t.Errorf("expected fix instruction, got:\n%s", got)
		}
	})

	t.Run("ExplicitDescAndRegoLayer", func(t *testing.T) {
		got := buildBlockFixUserPrompt(specsDir, ff, "CreateWorkflow", "explicit desc", "B", diags, layerRego)
		if !strings.Contains(got, "Feature: explicit desc") {
			t.Errorf("expected explicit desc, got:\n%s", got)
		}
	})

	t.Run("UnknownOpNoFeature", func(t *testing.T) {
		// op not in lookup → desc stays empty → no "Feature:" header, default layer.
		got := buildBlockFixUserPrompt(specsDir, ff, "NotAFeature", "", "B", diags, layerSSaC)
		if strings.Contains(got, "Feature:") {
			t.Errorf("did not expect Feature header for unknown op, got:\n%s", got)
		}
		if !strings.Contains(got, "OperationId: NotAFeature") {
			t.Errorf("expected operationId line, got:\n%s", got)
		}
	})
}

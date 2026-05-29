//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestBuildBlockUserPrompt — desc 유무/메시지 목록에 따른 prompt 구성 검증

package agent

import (
	"strings"
	"testing"
)

func TestBuildBlockUserPrompt(t *testing.T) {
	t.Run("WithDescAndMessages", func(t *testing.T) {
		got := buildBlockUserPrompt("the feature", "/v1/things", "things.ssac", "CreateThing", "BLOCK", []string{"e1: bad", "e2: worse"})
		for _, want := range []string{
			"Feature: the feature",
			"Path: /v1/things",
			"OperationId: CreateThing",
			"File: things.ssac",
			"Current block:\nBLOCK",
			"e1: bad",
			"e2: worse",
			"Fix ONLY this block.",
		} {
			if !strings.Contains(got, want) {
				t.Errorf("expected output to contain %q, got:\n%s", want, got)
			}
		}
	})

	t.Run("NoDescNoMessages", func(t *testing.T) {
		got := buildBlockUserPrompt("", "/ignored", "f.rego", "Op", "B", nil)
		if strings.Contains(got, "Feature:") {
			t.Errorf("did not expect Feature header when desc empty, got:\n%s", got)
		}
		if !strings.Contains(got, "OperationId: Op") {
			t.Errorf("expected operationId line, got:\n%s", got)
		}
		if !strings.Contains(got, "Validate errors:") {
			t.Errorf("expected errors header even with no messages, got:\n%s", got)
		}
	})
}

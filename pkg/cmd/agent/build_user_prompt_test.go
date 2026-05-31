//ff:func feature=agent type=test control=sequence
//ff:what TestBuildUserPrompt — desc 유무/메시지 목록에 따른 user prompt 구성 검증
package agent

import (
	"strings"
	"testing"
)

func TestBuildUserPrompt(t *testing.T) {
	t.Run("WithDescAndMessages", func(t *testing.T) {
		got := buildUserPrompt("the feature", "/v1/things", "things.ssac", "CONTENT", []string{"e1", "e2"})
		for _, want := range []string{
			"Feature: the feature",
			"Path: /v1/things",
			"Current file (things.ssac):\nCONTENT",
			"Validate errors:",
			"e1",
			"e2",
			"Fix the file.",
		} {
			if !strings.Contains(got, want) {
				t.Errorf("expected %q, got:\n%s", want, got)
			}
		}
	})

	t.Run("NoDescNoMessages", func(t *testing.T) {
		got := buildUserPrompt("", "/ignored", "f.rego", "C", nil)
		if strings.Contains(got, "Feature:") {
			t.Errorf("did not expect Feature header when desc empty, got:\n%s", got)
		}
		if !strings.Contains(got, "Current file (f.rego):\nC") {
			t.Errorf("expected current file section, got:\n%s", got)
		}
	})
}

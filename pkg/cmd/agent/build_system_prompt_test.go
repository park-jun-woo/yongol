//ff:func feature=agent type=test control=sequence
//ff:what TestBuildSystemPrompt — docs 섹션 매칭 시 포함, 미매칭/무문서 레이어는 예시만 포함 검증
package agent

import (
	"strings"
	"testing"
)

func TestBuildSystemPrompt(t *testing.T) {
	const base = "You fix yongol SSOT files."

	t.Run("WithMatchedDocSection", func(t *testing.T) {
		// layerSSaC has a doc file and the diag keywords match doc sections, so
		// the docSection branch is taken and appended before the example.
		diags := []string{"X-1 @auth resource", "D-2 BIGINT not null", "operationId mismatch"}
		got := buildSystemPrompt(layerSSaC, diags)
		if !strings.Contains(got, base) {
			t.Errorf("expected base prompt, got:\n%s", got)
		}
		if !strings.Contains(got, "Example for "+layerName(layerSSaC)) {
			t.Errorf("expected layer example header, got:\n%s", got)
		}
		// The matched doc section sits between the base and the example header.
		exampleHeader := "Example for "
		withSection := buildSystemPrompt(layerSSaC, diags)
		exampleOnly := buildSystemPrompt(layerUnknown, nil)
		if len(withSection) <= len(exampleOnly)+len(exampleHeader) {
			t.Errorf("expected doc section to add content; withSection=%d exampleOnly=%d", len(withSection), len(exampleOnly))
		}
	})

	t.Run("NoDocFileLayer", func(t *testing.T) {
		// layerUnknown maps to no doc file → searchDocs returns "" → example only.
		got := buildSystemPrompt(layerUnknown, []string{"@auth"})
		if !strings.Contains(got, base) {
			t.Errorf("expected base prompt, got:\n%s", got)
		}
		if !strings.Contains(got, "Example for "+layerName(layerUnknown)) {
			t.Errorf("expected example header, got:\n%s", got)
		}
	})
}

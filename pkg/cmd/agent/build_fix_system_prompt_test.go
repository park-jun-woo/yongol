//ff:func feature=agent type=test control=selection dimension=1
//ff:what TestBuildFixSystemPrompt — doc 파일 존재 시 doc 반환, 없는 레이어는 system prompt fallback 검증

package agent

import (
	"testing"

	"github.com/park-jun-woo/yongol/docs"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestBuildFixSystemPrompt(t *testing.T) {
	diags := []diagnostic.Diagnostic{{Message: "D-1: bad"}}

	t.Run("DocFileReturnedForKnownLayer", func(t *testing.T) {
		got := buildFixSystemPrompt(layerDDL, diags)
		want, err := docs.FS.ReadFile("ddl.md")
		if err != nil {
			t.Fatalf("read embedded ddl.md: %v", err)
		}
		if got != string(want) {
			t.Errorf("expected embedded ddl.md content for layerDDL")
		}
	})

	t.Run("DocFileReturnedForOpenAPI", func(t *testing.T) {
		got := buildFixSystemPrompt(layerOpenAPI, nil)
		want, err := docs.FS.ReadFile("openapi.md")
		if err != nil {
			t.Fatalf("read embedded openapi.md: %v", err)
		}
		if got != string(want) {
			t.Errorf("expected embedded openapi.md content for layerOpenAPI")
		}
	})

	t.Run("FallbackForUnknownLayer", func(t *testing.T) {
		// layerUnknown maps to no doc file → fallback to buildSystemPrompt.
		got := buildFixSystemPrompt(layerUnknown, diags)
		fallback := buildSystemPrompt(layerUnknown, diagMessages(diags))
		if got != fallback {
			t.Errorf("expected fallback system prompt for layerUnknown")
		}
		if got == "" {
			t.Error("fallback prompt should be non-empty")
		}
	})
}

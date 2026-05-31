//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestBuildGeneratePrompt — layer별(OpenAPI/Rego/Hurl/기본) 생성 프롬프트 구성 검증
package agent

import (
	"strings"
	"testing"
)

func TestBuildGeneratePrompt(t *testing.T) {
	const op = "CreateWorkflow"

	cases := []struct {
		name     string
		l        layer
		wantHas  string
		wantSkip string // marker that must NOT appear (cross-layer guard)
	}{
		{"OpenAPI", layerOpenAPI, "Generate a new OpenAPI path block", "Rego allow rule"},
		{"Rego", layerRego, "Generate a new Rego allow rule", "OpenAPI path block"},
		{"Hurl", layerHurl, "Generate a new Hurl request block", "Rego allow rule"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			assertGeneratePromptLayer(t, tc.l, op, tc.wantHas, tc.wantSkip)
		})
	}

	t.Run("DefaultLayerNoInstructions", func(t *testing.T) {
		// A layer without a generate template still emits the common header but
		// no layer-specific instruction block.
		got := buildGeneratePrompt(layerSSaC, op, "d", "/p", "S")
		if !strings.Contains(got, "Feature: d") {
			t.Errorf("expected feature header, got:\n%s", got)
		}
		for _, marker := range []string{"OpenAPI path block", "Rego allow rule", "Hurl request block"} {
			if strings.Contains(got, marker) {
				t.Errorf("did not expect %q for default layer, got:\n%s", marker, got)
			}
		}
	})
}

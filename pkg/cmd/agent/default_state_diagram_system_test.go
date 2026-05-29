//ff:func feature=agent type=test control=sequence
//ff:what TestDefaultStateDiagramSystem — fallback system prompt가 핵심 규칙·예시 포함 검증

package agent

import (
	"strings"
	"testing"
)

func TestDefaultStateDiagramSystem(t *testing.T) {
	out := defaultStateDiagramSystem()
	for _, want := range []string{
		"stateDiagram-v2",
		"[*] --> first_state",
		"operationId",
		"```mermaid",
		"ActivateWorkflow",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("prompt missing %q\n%s", want, out)
		}
	}
}

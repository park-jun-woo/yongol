//ff:func feature=statemachine type=parser control=sequence topic=states
//ff:what mermaid 블록이 파일 맨 앞에 위치할 때도 Transition.Line 이 올바른지 검증

package statemachine

import (
	"strings"
	"testing"
)

func TestParseTransitionLine_NoLeadingHeading(t *testing.T) {
	// mermaid block at the very start of the file.
	// Line 1: ```mermaid
	// Line 2: stateDiagram-v2
	// Line 3:     [*] --> a
	// Line 4:     a --> b: Go
	// Line 5: ```
	content := strings.Join([]string{
		"```mermaid",      // 1
		"stateDiagram-v2", // 2
		"    [*] --> a",   // 3
		"    a --> b: Go", // 4
		"```",             // 5
	}, "\n")

	d, diags := Parse("ab", content, "ab.md")
	if len(diags) > 0 {
		t.Fatalf("Parse diagnostics: %v", diags)
	}
	if len(d.Transitions) != 1 {
		t.Fatalf("Transitions count = %d, want 1", len(d.Transitions))
	}
	if d.Transitions[0].Line != 4 {
		t.Errorf("Transitions[0].Line = %d, want 4", d.Transitions[0].Line)
	}
}

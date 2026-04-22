//ff:func feature=statemachine type=parser control=iteration dimension=1 topic=states
//ff:what Transition.Line 이 mermaid 블록 내 실제 줄 번호로 채워지는지 검증

package statemachine

import (
	"strings"
	"testing"
)

func TestParseTransitionLine(t *testing.T) {
	// Carefully construct a fixture where each transition is on a known line.
	// Line 1: # Heading
	// Line 2: (empty)
	// Line 3: ```mermaid
	// Line 4: stateDiagram-v2
	// Line 5:     [*] --> draft
	// Line 6:     draft --> open: PublishGig
	// Line 7:     open --> closed: CloseGig
	// Line 8:     draft --> deleted: DeleteGig
	// Line 9: ```
	content := strings.Join([]string{
		"# Heading",                        // 1
		"",                                 // 2
		"```mermaid",                       // 3
		"stateDiagram-v2",                  // 4
		"    [*] --> draft",                // 5
		"    draft --> open: PublishGig",   // 6
		"    open --> closed: CloseGig",    // 7
		"    draft --> deleted: DeleteGig", // 8
		"```",                              // 9
		"",                                 // 10
	}, "\n")

	d, diags := Parse("gig", content, "gig.md")
	if len(diags) > 0 {
		t.Fatalf("Parse diagnostics: %v", diags)
	}
	if d == nil {
		t.Fatal("Parse returned nil diagram")
	}
	if len(d.Transitions) != 3 {
		t.Fatalf("Transitions count = %d, want 3", len(d.Transitions))
	}

	expected := []struct {
		event string
		line  int
	}{
		{"PublishGig", 6},
		{"CloseGig", 7},
		{"DeleteGig", 8},
	}
	for i, want := range expected {
		got := d.Transitions[i]
		if got.Event != want.event {
			t.Errorf("Transitions[%d].Event = %q, want %q", i, got.Event, want.event)
		}
		if got.Line != want.line {
			t.Errorf("Transitions[%d] (event=%s) Line = %d, want %d", i, got.Event, got.Line, want.line)
		}
	}
}

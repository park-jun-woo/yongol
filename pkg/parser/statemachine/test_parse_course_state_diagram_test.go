//ff:func feature=statemachine type=parser control=iteration dimension=1
//ff:what 코스 상태 다이어그램 파싱 검증 — ID, 초기 상태, 전이, 이벤트, ValidFromStates 확인

package statemachine

import (
	"testing"
)

func TestParseCourseStateDiagram(t *testing.T) {
	content := `# CourseState

` + "```mermaid" + `
stateDiagram-v2
    [*] --> unpublished
    unpublished --> published: PublishCourse
    published --> deleted: DeleteCourse
    unpublished --> deleted: DeleteCourse
` + "```" + `
`

	d, diags := Parse("course", content, "course.md")
	if len(diags) > 0 {
		t.Fatalf("Parse diagnostics: %v", diags)
	}

	if d.ID != "course" {
		t.Errorf("ID = %q, want %q", d.ID, "course")
	}

	if d.InitialState != "unpublished" {
		t.Errorf("InitialState = %q, want %q", d.InitialState, "unpublished")
	}

	if len(d.Transitions) != 3 {
		t.Fatalf("Transitions count = %d, want 3", len(d.Transitions))
	}

	// Check transitions.
	expected := []Transition{
		{From: "unpublished", To: "published", Event: "PublishCourse"},
		{From: "published", To: "deleted", Event: "DeleteCourse"},
		{From: "unpublished", To: "deleted", Event: "DeleteCourse"},
	}
	for i, want := range expected {
		got := d.Transitions[i]
		if got.From != want.From || got.To != want.To || got.Event != want.Event {
			t.Errorf("Transition[%d] = %+v, want %+v", i, got, want)
		}
	}

	// Check Events().
	events := d.Events()
	if len(events) != 2 {
		t.Errorf("Events count = %d, want 2", len(events))
	}

	// Check ValidFromStates.
	fromStates := d.ValidFromStates("DeleteCourse")
	if len(fromStates) != 2 {
		t.Errorf("ValidFromStates(DeleteCourse) count = %d, want 2", len(fromStates))
	}
}

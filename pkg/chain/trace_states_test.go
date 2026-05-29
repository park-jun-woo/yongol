//ff:func feature=chain type=test control=iteration dimension=2
//ff:what traceStates 가 @state sequence 를 매칭 stateDiagram 과 연결하고 미참조 시 nil 을 반환하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func TestTraceStates(t *testing.T) {
	specsDir := t.TempDir()
	statesDir := filepath.Join(specsDir, "states")
	if err := os.MkdirAll(statesDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(statesDir, "reservation.md"), []byte("stateDiagram-v2\n  Pending --> Cancelled: cancel\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	diagrams := []*statemachine.StateDiagram{
		{ID: "reservation"},
		{ID: "unrelated"},
	}

	sf := &ssac.ServiceFunc{
		Name: "CancelReservation",
		Sequences: []ssac.Sequence{
			{Type: "state", DiagramID: "reservation", Transition: "cancel"},
			{Type: "get", Model: "Reservation.FindByID"}, // ignored
		},
	}

	links := traceStates(sf, diagrams, specsDir)
	if len(links) != 1 {
		t.Fatalf("expected 1 state link, got %d: %+v", len(links), links)
	}
	if links[0].Kind != "StateDiag" || links[0].File != "states/reservation.md" {
		t.Errorf("link fields: %+v", links[0])
	}
	if links[0].Summary != "diagram: reservation -> cancel" {
		t.Errorf("summary: %q", links[0].Summary)
	}

	// No @state sequences → nil.
	sfNone := &ssac.ServiceFunc{Name: "X", Sequences: []ssac.Sequence{{Type: "get", Model: "Y.Z"}}}
	if traceStates(sfNone, diagrams, specsDir) != nil {
		t.Error("expected nil when no @state sequences")
	}
}

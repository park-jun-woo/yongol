//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildState_ZeroCov — @state 전이 검증 (Symbol fallback 포함)
package ssac

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestBuildState_ZeroCov(t *testing.T) {
	g := &methodGen{
		FuncName:      "CancelReservation",
		ModulePath:    "example.com/app",
		DiagramSymbol: map[string]string{"reservation": "Reservation"},
	}
	seq := ssacparser.Sequence{
		Type:       "state",
		DiagramID:  "reservation",
		Inputs:     map[string]string{"status": "reservation.Status"},
		Transition: "cancel",
	}
	lines, imports := g.buildState(seq)
	body := strings.Join(lines, "\n")
	if !strings.Contains(body, "statemachine.ReservationCanTransition(") {
		t.Fatalf("expected statemachine call, got:\n%s", body)
	}
	if !strings.Contains(body, `"cancel"`) {
		t.Fatalf("expected transition literal, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(imports, " "), "internal/statemachine") {
		t.Fatalf("expected statemachine import, got %v", imports)
	}

	// Fallback: DiagramSymbol missing → uses DiagramID.
	g2 := &methodGen{FuncName: "X", ModulePath: "m", DiagramSymbol: map[string]string{}}
	seq2 := ssacparser.Sequence{Type: "state", DiagramID: "order", Inputs: map[string]string{"s": "o.S"}, Transition: "ship"}
	body2 := strings.Join(func() []string { l, _ := g2.buildState(seq2); return l }(), "\n")
	if !strings.Contains(body2, "statemachine.orderCanTransition(") {
		t.Fatalf("expected fallback to DiagramID, got:\n%s", body2)
	}
}

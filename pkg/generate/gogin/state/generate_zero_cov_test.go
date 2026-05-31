//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildTransitionMap_ZeroCov — 초기 전이 제외 + 중첩 맵 구성
package state

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_ZeroCov(t *testing.T) {
	// Empty → no-op
	if err := Generate(&yongol.Fullstack{}, t.TempDir()); err != nil {
		t.Fatalf("empty Generate error: %v", err)
	}
	// With diagram → files written under backend/internal/statemachine
	artifacts := t.TempDir()
	fs := &yongol.Fullstack{
		StateDiagrams: []*statemachine.StateDiagram{
			{
				ID:     "course",
				Symbol: "Course",
				Transitions: []statemachine.Transition{
					{From: "[*]", To: "draft", Event: "create"},
					{From: "draft", To: "review", Event: "submit"},
				},
			},
		},
	}
	if err := Generate(fs, artifacts); err != nil {
		t.Fatalf("Generate error: %v", err)
	}
	dir := filepath.Join(artifacts, "backend", "internal", "statemachine")
	if _, err := os.Stat(filepath.Join(dir, "course.go")); err != nil {
		t.Errorf("expected course.go generated: %v", err)
	}
}

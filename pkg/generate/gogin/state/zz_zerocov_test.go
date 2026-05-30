package state

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBuildTransitionMap_ZeroCov — 초기 전이 제외 + 중첩 맵 구성

func TestBuildTransitionMap_ZeroCov(t *testing.T) {
	d := &statemachine.StateDiagram{
		Transitions: []statemachine.Transition{
			{From: "[*]", To: "draft", Event: "create"},
			{From: "draft", To: "review", Event: "submit"},
			{From: "draft", To: "archived", Event: "archive"},
			{From: "review", To: "published", Event: "approve"},
		},
	}
	m := buildTransitionMap(d)
	if _, ok := m["[*]"]; ok {
		t.Error("initial transition should be excluded")
	}
	if m["draft"]["submit"] != "review" {
		t.Errorf("draft/submit = %q", m["draft"]["submit"])
	}
	if len(m["draft"]) != 2 {
		t.Errorf("draft events = %d", len(m["draft"]))
	}
}

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestSortedKeys_ZeroCov — 정렬된 키 반환

func TestSortedKeys_ZeroCov(t *testing.T) {
	m := map[string]int{"c": 1, "a": 2, "b": 3}
	got := sortedKeys(m)
	if strings.Join(got, ",") != "a,b,c" {
		t.Errorf("sortedKeys = %v", got)
	}
}

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRenderTransitionEntries_ZeroCov — 단일/다중 이벤트 렌더링 분기

func TestRenderTransitionEntries_ZeroCov(t *testing.T) {
	var b strings.Builder
	transMap := map[string]map[string]string{
		"draft":  {"submit": "review"},                    // single event branch
		"review": {"approve": "published", "reject": "draft"}, // multi event branch
	}
	renderTransitionEntries(&b, transMap)
	out := b.String()
	if !strings.Contains(out, `"draft": {"submit": "review"}`) {
		t.Errorf("single-event line missing:\n%s", out)
	}
	if !strings.Contains(out, `"review": {`) || !strings.Contains(out, `"approve": "published"`) {
		t.Errorf("multi-event block missing:\n%s", out)
	}
}

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRenderStateFile_ZeroCov — Transitions 변수 source 생성

func TestRenderStateFile_ZeroCov(t *testing.T) {
	src := renderStateFile("course", "Course", map[string]map[string]string{
		"draft": {"submit": "review"},
	})
	for _, want := range []string{"package statemachine", "var CourseTransitions", "states/course.md"} {
		if !strings.Contains(src, want) {
			t.Errorf("renderStateFile missing %q", want)
		}
	}
}

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRenderCanTransitionFile_ZeroCov — CanTransition 가드 source 생성

func TestRenderCanTransitionFile_ZeroCov(t *testing.T) {
	src := renderCanTransitionFile("course", "Course")
	for _, want := range []string{"package statemachine", "func CourseCanTransition", "CourseTransitions"} {
		if !strings.Contains(src, want) {
			t.Errorf("renderCanTransitionFile missing %q", want)
		}
	}
}

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRenderNextStateFile_ZeroCov — NextState 접근자 source 생성

func TestRenderNextStateFile_ZeroCov(t *testing.T) {
	src := renderNextStateFile("course", "Course")
	for _, want := range []string{"package statemachine", "func CourseNextState", "CourseTransitions"} {
		if !strings.Contains(src, want) {
			t.Errorf("renderNextStateFile missing %q", want)
		}
	}
}

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWriteStateFile_ZeroCov — 3종 파일 디스크 기록

func TestWriteStateFile_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	transMap := map[string]map[string]string{"draft": {"submit": "review"}}
	if err := writeStateFile(dir, "course", "Course", transMap); err != nil {
		t.Fatalf("writeStateFile error: %v", err)
	}
	for _, name := range []string{"course.go", "course_can_transition.go", "course_next_state.go"} {
		if _, err := os.Stat(filepath.Join(dir, name)); err != nil {
			t.Errorf("expected file %s: %v", name, err)
		}
	}
}

//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerate_ZeroCov — empty(skip) + StateDiagram → 파일 생성

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

//ff:func feature=agent type=test control=sequence
//ff:what TestFixFileGroup — stall 추적(증가/리셋/3회skip) + fixFile 성공/실패 + ruleID hint 분기 검증

package agent

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func writeManifest(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("metadata: {}"), 0o644); err != nil {
		t.Fatal(err)
	}
	return dir
}

func TestFixFileGroupSuccess(t *testing.T) {
	dir := writeManifest(t)
	llm := func(b, m, s, u string) (string, error) { return "fixed: yes\n", nil }
	g := fileGroup{
		relFile: "manifest.yaml",
		diags:   []diagnostic.Diagnostic{{Message: "[S-01] broken", Level: diagnostic.LevelError}},
	}
	var out bytes.Buffer
	stalled := map[string]*stallTracker{}
	ok := fixFileGroup(dir, &features.FeaturesFile{}, g, llm, Config{}, &out, stalled)
	if !ok {
		t.Fatalf("expected success, out=%q", out.String())
	}
	// ruleID hint should be appended.
	if !strings.Contains(out.String(), "fixed: manifest.yaml (S-01)") {
		t.Errorf("expected ruleID hint, got: %q", out.String())
	}
}

func TestFixFileGroupSuccessNoRuleID(t *testing.T) {
	dir := writeManifest(t)
	llm := func(b, m, s, u string) (string, error) { return "fixed\n", nil }
	g := fileGroup{
		relFile: "manifest.yaml",
		diags:   []diagnostic.Diagnostic{{Message: "no rule id here", Level: diagnostic.LevelError}},
	}
	var out bytes.Buffer
	ok := fixFileGroup(dir, &features.FeaturesFile{}, g, llm, Config{}, &out, map[string]*stallTracker{})
	if !ok {
		t.Fatalf("expected success, out=%q", out.String())
	}
	if strings.Contains(out.String(), "(") {
		t.Errorf("expected no ruleID hint, got: %q", out.String())
	}
}

func TestFixFileGroupNoDiags(t *testing.T) {
	// Empty diags exercises the len(g.diags) == 0 branch (no hint lookup).
	dir := writeManifest(t)
	llm := func(b, m, s, u string) (string, error) { return "ok\n", nil }
	g := fileGroup{relFile: "manifest.yaml"}
	var out bytes.Buffer
	if !fixFileGroup(dir, &features.FeaturesFile{}, g, llm, Config{}, &out, map[string]*stallTracker{}) {
		t.Fatalf("expected success, out=%q", out.String())
	}
}

func TestFixFileGroupFixError(t *testing.T) {
	dir := writeManifest(t)
	llm := func(b, m, s, u string) (string, error) { return "", errors.New("llm down") }
	g := fileGroup{
		relFile: "manifest.yaml",
		diags:   []diagnostic.Diagnostic{{Message: "x", Level: diagnostic.LevelError}},
	}
	var out bytes.Buffer
	if fixFileGroup(dir, &features.FeaturesFile{}, g, llm, Config{}, &out, map[string]*stallTracker{}) {
		t.Fatal("expected failure when fixFile errors")
	}
	if !strings.Contains(out.String(), "skipped") {
		t.Errorf("expected skipped message, got: %q", out.String())
	}
}

func TestFixFileGroupStall(t *testing.T) {
	dir := writeManifest(t)
	llm := func(b, m, s, u string) (string, error) { return "ok\n", nil }
	g := fileGroup{
		relFile: "manifest.yaml",
		diags:   []diagnostic.Diagnostic{{Message: "same error", Level: diagnostic.LevelError}},
	}
	stalled := map[string]*stallTracker{}

	// Round 1 & 2: same messages, tracker.count increments to 2 — still fixes.
	var out bytes.Buffer
	if !fixFileGroup(dir, &features.FeaturesFile{}, g, llm, Config{}, &out, stalled) {
		t.Fatal("round 1 should succeed")
	}
	if !fixFileGroup(dir, &features.FeaturesFile{}, g, llm, Config{}, &out, stalled) {
		t.Fatal("round 2 should succeed")
	}
	// Round 3: count reaches 3 → stalled, returns false.
	out.Reset()
	if fixFileGroup(dir, &features.FeaturesFile{}, g, llm, Config{}, &out, stalled) {
		t.Fatal("round 3 should be stalled")
	}
	if !strings.Contains(out.String(), "stalled") {
		t.Errorf("expected stalled message, got: %q", out.String())
	}
}

func TestFixFileGroupStallReset(t *testing.T) {
	// Changing the diagnostic messages resets the stall counter (else branch).
	dir := writeManifest(t)
	llm := func(b, m, s, u string) (string, error) { return "ok\n", nil }
	stalled := map[string]*stallTracker{}
	var out bytes.Buffer

	g1 := fileGroup{relFile: "manifest.yaml", diags: []diagnostic.Diagnostic{{Message: "err A"}}}
	g2 := fileGroup{relFile: "manifest.yaml", diags: []diagnostic.Diagnostic{{Message: "err B"}}}

	fixFileGroup(dir, &features.FeaturesFile{}, g1, llm, Config{}, &out, stalled)
	fixFileGroup(dir, &features.FeaturesFile{}, g2, llm, Config{}, &out, stalled) // resets count to 1
	if stalled["manifest.yaml"].count != 1 {
		t.Errorf("expected count reset to 1 after message change, got %d", stalled["manifest.yaml"].count)
	}
}

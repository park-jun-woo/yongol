//ff:func feature=agent type=test control=sequence
//ff:what TestFixFileGroup — stall 추적(증가/리셋/3회skip) + fixFile 성공/실패 + ruleID hint 분기 검증
package agent

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

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

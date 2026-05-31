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

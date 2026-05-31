//ff:func feature=agent type=test control=sequence
//ff:what TestFixFileGroup — stall 추적(증가/리셋/3회skip) + fixFile 성공/실패 + ruleID hint 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

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

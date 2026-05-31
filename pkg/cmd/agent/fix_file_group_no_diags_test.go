//ff:func feature=agent type=test control=sequence
//ff:what TestFixFileGroup — stall 추적(증가/리셋/3회skip) + fixFile 성공/실패 + ruleID hint 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

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

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

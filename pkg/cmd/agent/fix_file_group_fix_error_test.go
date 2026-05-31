//ff:func feature=agent type=test control=sequence
//ff:what TestFixFileGroup — stall 추적(증가/리셋/3회skip) + fixFile 성공/실패 + ruleID hint 분기 검증
package agent

import (
	"bytes"
	"errors"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

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

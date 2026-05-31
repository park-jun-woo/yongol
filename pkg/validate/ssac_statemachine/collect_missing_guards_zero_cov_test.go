//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestStatemachineBatch_ZeroCov — ssac_statemachine 빌더/수집기 헬퍼 직접 커버
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectMissingGuards_ZeroCov(t *testing.T) {
	funcByName := map[string]ssac.ServiceFunc{
		"activate": {Name: "activate", FileName: "a.ssac", Line: 2},
		"close":    {Name: "close"},
	}
	guards := map[string]bool{"close": true}
	// "activate": has func, not a guard → diag; "close": guard → skip; "missing": no func → skip
	diags := collectMissingGuards("wf", []string{"activate", "close", "missing"}, funcByName, guards)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	if diags[0].File != "a.ssac" {
		t.Errorf("file = %q", diags[0].File)
	}
	// func with empty FileName → synthesized
	fb := map[string]ssac.ServiceFunc{"go": {Name: "go"}}
	d := collectMissingGuards("wf", []string{"go"}, fb, map[string]bool{})
	if d[0].File != "ssac/go.ssac" {
		t.Errorf("synth file = %q", d[0].File)
	}
}

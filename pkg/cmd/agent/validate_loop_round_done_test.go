//ff:func feature=agent type=test control=sequence
//ff:what TestValidateLoopRound — 0 fixable 에러 done=true / runValidate 에러 전파 분기 검증
package agent

import (
	"bytes"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestValidateLoopRoundDone(t *testing.T) {
	// Empty valid specs -> 0 fixable errors -> done=true.
	var out bytes.Buffer
	ff := &features.FeaturesFile{}
	stalled := map[string]*stallTracker{}
	totalFixed := 0
	dir := t.TempDir()
	res, err := validateLoopRound(dir, ff, nil, Config{}, &out, dir, 1, stalled, &totalFixed)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !res.done {
		t.Fatalf("expected done=true on clean specs, got %+v", res)
	}
}

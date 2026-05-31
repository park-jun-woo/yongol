//ff:func feature=agent type=test control=sequence
//ff:what TestValidateLoopRound — 0 fixable 에러 done=true / runValidate 에러 전파 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestValidateLoopRoundValidateError(t *testing.T) {
	// A file (not a directory) makes runValidate fail -> error propagated.
	file := filepath.Join(t.TempDir(), "f.txt")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	ff := &features.FeaturesFile{}
	stalled := map[string]*stallTracker{}
	totalFixed := 0
	if _, err := validateLoopRound(file, ff, nil, Config{}, &out, file, 1, stalled, &totalFixed); err == nil {
		t.Fatal("expected error from runValidate")
	}
}

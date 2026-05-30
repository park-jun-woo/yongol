//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestValidateLoop — 빈 specs 0 에러 조기종료 성공 / runValidate 에러 전파 분기 검증

package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestValidateLoopCleanConverges(t *testing.T) {
	// Empty (valid) specs dir -> first round has 0 fixable errors -> done, nil.
	var out, errOut bytes.Buffer
	ff := &features.FeaturesFile{}
	if err := validateLoop(t.TempDir(), ff, nil, Config{}, &out, &errOut, 3); err != nil {
		t.Fatalf("clean specs: unexpected error: %v", err)
	}
}

func TestValidateLoopValidateError(t *testing.T) {
	// Passing a file (not a directory) makes runValidate (DetectSSOTs) fail,
	// propagated out of the round.
	file := filepath.Join(t.TempDir(), "f.txt")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	var out, errOut bytes.Buffer
	ff := &features.FeaturesFile{}
	if err := validateLoop(file, ff, nil, Config{}, &out, &errOut, 2); err == nil {
		t.Fatal("expected error from runValidate on non-directory specs")
	}
}

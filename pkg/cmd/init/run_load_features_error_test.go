//ff:func feature=cli-init type=test control=sequence
//ff:what TestRunExtraBranches — loadFeatures 에러 / Dir 기본값 / Module 자동탐지 warning 분기 검증
package cliinit

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestRun_LoadFeaturesError(t *testing.T) {
	// FeaturesPath points at a missing file -> loadFeatures fails before writes.
	var out, errOut bytes.Buffer
	err := Run(&out, &errOut, Options{
		ProjectID:    "Myapp",
		FeaturesPath: filepath.Join(t.TempDir(), "nope.yaml"),
		Dir:          filepath.Join(t.TempDir(), "x"),
		Module:       "github.com/test/x",
	})
	if err == nil {
		t.Fatal("want error for unreadable features.yaml")
	}
}

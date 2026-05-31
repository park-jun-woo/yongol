//ff:func feature=cli-init type=test control=sequence
//ff:what TestDetectModuleBranches — GITHUB_USER/GH_USER 우선순위·정상 모듈·비안전문자 placeholder 분기 검증
package cliinit

import (
	"testing"
)

func TestDetectModule_GHUserFallback(t *testing.T) {
	t.Setenv("GITHUB_USER", "")
	t.Setenv("GH_USER", "Carol")
	module, warning := DetectModule("App")
	if module != "github.com/carol/App" {
		t.Errorf("module = %q, want github.com/carol/App", module)
	}
	if warning != "" {
		t.Errorf("warning should be empty, got %q", warning)
	}
}

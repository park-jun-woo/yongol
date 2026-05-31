//ff:func feature=cli-init type=test control=sequence
//ff:what TestDetectModuleBranches — GITHUB_USER/GH_USER 우선순위·정상 모듈·비안전문자 placeholder 분기 검증
package cliinit

import (
	"strings"
	"testing"
)

func TestDetectModule_UnsafeUserPlaceholder(t *testing.T) {
	// A user name with no GitHub-safe characters normalizes to "" -> placeholder.
	t.Setenv("GITHUB_USER", "***")
	t.Setenv("GH_USER", "")
	module, warning := DetectModule("App")
	if !strings.HasPrefix(module, "github.com/REPLACE_ME/") {
		t.Errorf("module = %q, want placeholder prefix", module)
	}
	if warning == "" {
		t.Error("warning should be set for unsafe user name")
	}
}

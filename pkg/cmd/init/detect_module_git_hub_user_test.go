//ff:func feature=cli-init type=test control=sequence
//ff:what TestDetectModuleBranches — GITHUB_USER/GH_USER 우선순위·정상 모듈·비안전문자 placeholder 분기 검증
package cliinit

import (
	"testing"
)

func TestDetectModule_GitHubUser(t *testing.T) {
	t.Setenv("GITHUB_USER", "Alice")
	t.Setenv("GH_USER", "bob")
	module, warning := DetectModule("MyApp")
	if module != "github.com/alice/MyApp" {
		t.Errorf("module = %q, want github.com/alice/MyApp", module)
	}
	if warning != "" {
		t.Errorf("warning should be empty, got %q", warning)
	}
}

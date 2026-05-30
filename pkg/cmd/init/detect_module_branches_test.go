//ff:func feature=cli-init type=test control=sequence
//ff:what TestDetectModuleBranches — GITHUB_USER/GH_USER 우선순위·정상 모듈·비안전문자 placeholder 분기 검증

package cliinit

import (
	"strings"
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

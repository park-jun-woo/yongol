//ff:type feature=cli-init type=test
//ff:what test — DetectModule honors GITHUB_USER and falls back to placeholder

package cliinit

import (
	"strings"
	"testing"
)

func TestDetectModuleWithGitHubUserEnv(t *testing.T) {
	t.Setenv("GITHUB_USER", "alice")
	t.Setenv("GH_USER", "")
	module, warning := DetectModule("MyApp")
	if module != "github.com/alice/MyApp" {
		t.Errorf("module = %q, want github.com/alice/MyApp", module)
	}
	if warning != "" {
		t.Errorf("warning = %q, want empty", warning)
	}
}

func TestDetectModuleFallsBackToPlaceholder(t *testing.T) {
	t.Setenv("GITHUB_USER", "")
	t.Setenv("GH_USER", "")
	// Force git user.name probe to fail by pointing HOME at a directory with
	// no git config. We cannot disable git outright, but when git finds no
	// user.name it exits non-zero.
	t.Setenv("HOME", t.TempDir())
	t.Setenv("XDG_CONFIG_HOME", t.TempDir())
	t.Setenv("GIT_CONFIG_GLOBAL", "/dev/null")
	module, warning := DetectModule("MyApp")
	if !strings.HasPrefix(module, "github.com/REPLACE_ME/") {
		t.Errorf("module = %q, want prefix github.com/REPLACE_ME/", module)
	}
	if warning == "" {
		t.Errorf("warning should be set when falling back to placeholder")
	}
}

func TestNormalizeGitHubUser(t *testing.T) {
	cases := []struct{ in, want string }{
		{"Alice", "alice"},
		{"Park Jun Woo", "parkjunwoo"},
		{"park-jun-woo", "park-jun-woo"},
		{"alice.bob", "alicebob"},
		{"alice@corp", "alicecorp"},
	}
	for _, tc := range cases {
		got := normalizeGitHubUser(tc.in)
		if got != tc.want {
			t.Errorf("normalizeGitHubUser(%q) = %q, want %q", tc.in, got, tc.want)
		}
	}
}

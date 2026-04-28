//ff:func feature=cli-init type=test control=sequence
//ff:what TestDetectModuleFallsBackToPlaceholder — DetectModule uses placeholder when no user detectable

package cliinit

import (
	"strings"
	"testing"
)

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

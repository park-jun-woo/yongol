//ff:func feature=cli-init type=test control=sequence
//ff:what TestDetectModuleWithGitHubUserEnv — DetectModule honors GITHUB_USER env var

package cliinit

import "testing"

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

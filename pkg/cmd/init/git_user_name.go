//ff:func feature=cli-init type=util control=sequence
//ff:what gitUserName — read `git config --global user.name` (empty on any failure)

package cliinit

import (
	"os/exec"
	"strings"
)

// gitUserName reads `git config --global user.name`. Any git failure (git not
// installed, config unset, non-zero exit) is treated as "unknown".
func gitUserName() (string, bool) {
	cmd := exec.Command("git", "config", "--global", "user.name")
	out, err := cmd.Output()
	if err != nil {
		return "", false
	}
	name := strings.TrimSpace(string(out))
	if name == "" {
		return "", false
	}
	return name, true
}

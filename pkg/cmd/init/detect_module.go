//ff:func feature=cli-init type=util control=sequence
//ff:what detectModule — heuristics for inferring a Go module path when the user omits --module

package cliinit

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// ModulePlaceholder is substituted when neither env vars nor git config can
// identify a GitHub user. Surfacing a deliberately wrong path is safer than
// making up a user name — the warning printed by DetectModule tells the user
// to fix it explicitly.
const ModulePlaceholder = "github.com/REPLACE_ME"

// DetectModule returns the Go module path to use in manifest.backend.module
// along with a warning string when the result relies on a placeholder. The
// caller is expected to print the warning on stderr (never on stdout) so
// tooling parsing the command output is not disturbed.
//
// Resolution order:
//  1. GITHUB_USER environment variable
//  2. GH_USER environment variable (gh CLI convention)
//  3. `git config --global user.name` (space-stripped, lowercased)
//  4. placeholder "github.com/REPLACE_ME"
func DetectModule(projectID string) (module string, warning string) {
	user := strings.TrimSpace(os.Getenv("GITHUB_USER"))
	if user == "" {
		user = strings.TrimSpace(os.Getenv("GH_USER"))
	}
	if user == "" {
		if name, ok := gitUserName(); ok {
			user = name
		}
	}
	if user == "" {
		return fmt.Sprintf("%s/%s", ModulePlaceholder, projectID),
			"no GITHUB_USER / GH_USER / git user.name detected — using placeholder; override with --module"
	}
	normalized := normalizeGitHubUser(user)
	if normalized == "" {
		return fmt.Sprintf("%s/%s", ModulePlaceholder, projectID),
			"detected user name contains no GitHub-safe characters — using placeholder; override with --module"
	}
	return fmt.Sprintf("github.com/%s/%s", normalized, projectID), ""
}

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

// githubUserAllowed matches characters that are valid in a GitHub username
// slug — letters, digits, hyphens. Everything else is stripped.
var githubUserAllowed = regexp.MustCompile(`[^A-Za-z0-9-]+`)

// normalizeGitHubUser lowercases and strips disallowed characters so "Park Jun
// Woo" becomes "parkjunwoo" and "Park-Jun Woo" becomes "park-junwoo". The
// result is a best-effort guess — the user is expected to pass --module
// explicitly when the inference is wrong.
func normalizeGitHubUser(name string) string {
	lower := strings.ToLower(name)
	// Spaces and separators collapse into empty string (not hyphen) to avoid
	// inventing a hyphen the user never actually uses as a GitHub handle.
	lower = strings.ReplaceAll(lower, " ", "")
	lower = githubUserAllowed.ReplaceAllString(lower, "")
	lower = strings.Trim(lower, "-")
	return lower
}

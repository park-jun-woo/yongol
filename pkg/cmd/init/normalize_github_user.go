//ff:func feature=cli-init type=util control=sequence
//ff:what normalizeGitHubUser — lowercase and strip disallowed characters from a git user.name

package cliinit

import (
	"regexp"
	"strings"
)

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

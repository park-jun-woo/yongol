//ff:func feature=gen-hurl type=util control=sequence
//ff:what isUserHurlName — user-authored hurl 파일명 판정 (smoke.hurl 제외)

package hurl

import (
	"strings"
)

// isUserHurlName reports whether name is a user-authored hurl file
// that should be mirrored. smoke.hurl is explicitly excluded so the
// auto-generator never clashes with a stray user file of the same
// name.
func isUserHurlName(name string) bool {
	if !strings.HasSuffix(name, ".hurl") {
		return false
	}
	if name == "smoke.hurl" {
		return false
	}
	return strings.HasPrefix(name, "scenario-") || strings.HasPrefix(name, "invariant-")
}

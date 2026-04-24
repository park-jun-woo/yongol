//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what isUserCreateWithPasswordHash — @post <Model>.Create Inputs 에 PasswordHash 포함 여부

package hurl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// isUserCreateWithPasswordHash reports whether seq is a @post <Model>.Create
// sequence whose Inputs include a PasswordHash-cased key. Extracted so the
// caller (hasUserCreatePost) stays at iteration dimension=1 (single loop
// at depth 1).
func isUserCreateWithPasswordHash(seq ssac.Sequence) bool {
	if seq.Type != ssac.SeqPost {
		return false
	}
	if !strings.HasSuffix(seq.Model, ".Create") {
		return false
	}
	for k := range seq.Inputs {
		if strings.EqualFold(k, "PasswordHash") {
			return true
		}
	}
	return false
}

//ff:func feature=report type=util control=sequence topic=json
//ff:what tryAbsRelativeFile — 절대경로 기반으로 specsDir 상대경로 재시도 (실패 시 "" 반환)

package json

import (
	"path/filepath"
	"strings"
)

// tryAbsRelativeFile returns file expressed relative to absSpecs using
// absolute-versus-absolute comparison, or "" when the rebase isn't possible
// or escapes the root ("..").
func tryAbsRelativeFile(file, absSpecs string) string {
	if absSpecs == "" {
		return ""
	}
	absFile, err := filepath.Abs(file)
	if err != nil {
		return ""
	}
	rel, err := filepath.Rel(absSpecs, absFile)
	if err != nil || strings.HasPrefix(rel, "..") {
		return ""
	}
	return filepath.ToSlash(rel)
}

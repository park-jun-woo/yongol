//ff:func feature=report type=util control=sequence topic=sarif
//ff:what tryAbsRelativeURI — 절대경로 기반으로 specsDir 상대경로 재시도 (실패 시 "" 반환)

package sarif

import (
	"path/filepath"
	"strings"
)

// tryAbsRelativeURI attempts an absolute-versus-absolute rebase of file
// against absSpecs. Returns the forward-slash relative path when possible,
// or "" when the rebase fails or escapes the root.
func tryAbsRelativeURI(file, absSpecs string) string {
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

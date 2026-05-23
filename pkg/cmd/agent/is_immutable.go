//ff:func feature=agent type=helper control=sequence
//ff:what isImmutable — 파일 경로가 변경 불가 소스(features.yaml, .hurl, .yongol)인지 판별

package agent

import (
	"path/filepath"
	"strings"
)

// isImmutable returns true if the file path refers to an immutable source.
func isImmutable(file string) bool {
	base := filepath.Base(file)
	if base == "features.yaml" {
		return true
	}
	if base == ".yongol" {
		return true
	}
	if strings.HasSuffix(file, ".hurl") {
		return true
	}
	return false
}

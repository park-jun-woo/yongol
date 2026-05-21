//ff:func feature=agent type=helper control=pure
//ff:what filterImmutable — immutable 파일(features.yaml, .hurl, .yongol)의 diagnostic 제외

package agent

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// filterImmutable removes diagnostics whose File refers to an immutable source:
//   - features.yaml
//   - tests/*.hurl or any file ending in .hurl
//   - .yongol
func filterImmutable(diags []diagnostic.Diagnostic) []diagnostic.Diagnostic {
	out := make([]diagnostic.Diagnostic, 0, len(diags))
	for _, d := range diags {
		if isImmutable(d.File) {
			continue
		}
		out = append(out, d)
	}
	return out
}

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

// countImmutable returns how many diagnostics target immutable files.
func countImmutable(diags []diagnostic.Diagnostic) int {
	n := 0
	for _, d := range diags {
		if d.Level == diagnostic.LevelError && isImmutable(d.File) {
			n++
		}
	}
	return n
}

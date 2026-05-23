//ff:func feature=agent type=helper control=sequence
//ff:what rebaseFile — 절대 경로를 specsDir 상대 경로로 변환

package agent

import "path/filepath"

// rebaseFile converts an absolute file path to a specsDir-relative path.
// If already relative or conversion fails, returns the original.
func rebaseFile(file, absSpecs string) string {
	if !filepath.IsAbs(file) {
		return file
	}
	rel, err := filepath.Rel(absSpecs, file)
	if err != nil {
		return file
	}
	return filepath.ToSlash(rel)
}

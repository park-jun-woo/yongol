//ff:func feature=orchestrator type=util control=sequence
//ff:what dirPresence — 디렉토리 존재 + 파일 수 → SSOTPresence 매핑

package yongol

import "os"

// dirPresence classifies an SSOT directory into one of the three presence
// states based on whether the directory itself exists and whether any matching
// content files were found. Callers prefilter content via glob patterns and
// pass the resulting count.
func dirPresence(dir string, fileCount int) SSOTPresence {
	fi, err := os.Stat(dir)
	dirExists := err == nil && fi.IsDir()
	if fileCount > 0 {
		return SSOTPopulated
	}
	if dirExists {
		return SSOTDeclared
	}
	return SSOTAbsent
}

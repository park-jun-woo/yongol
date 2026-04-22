//ff:func feature=orchestrator type=util control=sequence
//ff:what trySSaCPkgPath — yongol root 형제 경로의 ssac/pkg 디렉토리 존재 여부 확인
package yongol

import (
	"os"
	"path/filepath"
)

// trySSaCPkgPath returns `<parent>/ssac/pkg` if it exists as a directory,
// otherwise returns "".
func trySSaCPkgPath(yongolRoot string) string {
	ssacPkg := filepath.Join(filepath.Dir(yongolRoot), "ssac", "pkg")
	fi, err := os.Stat(ssacPkg)
	if err != nil || !fi.IsDir() {
		return ""
	}
	return ssacPkg
}

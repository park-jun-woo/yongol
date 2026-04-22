//ff:func feature=orchestrator type=util control=sequence
//ff:what 디렉토리가 yongol go.mod + pkg/ 를 가진 루트인지 판별
package yongol

import (
	"os"
	"path/filepath"
	"strings"
)

// isYongolRoot returns true if dir contains a yongol go.mod and a pkg/ subdirectory.
func isYongolRoot(dir string) bool {
	data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil || !strings.Contains(string(data), "github.com/park-jun-woo/yongol") {
		return false
	}
	fi, err := os.Stat(filepath.Join(dir, "pkg"))
	return err == nil && fi.IsDir()
}

//ff:func feature=orchestrator type=util control=sequence
//ff:what isDir — 경로가 존재하고 디렉토리인지 확인

package yongol

import "os"

// isDir returns true if p exists and is a directory.
func isDir(p string) bool {
	fi, err := os.Stat(p)
	return err == nil && fi.IsDir()
}

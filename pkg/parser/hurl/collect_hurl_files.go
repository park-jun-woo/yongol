//ff:func feature=orchestrator type=util control=iteration dimension=1
//ff:what 디렉토리에서 .hurl 파일 경로 목록을 수집
package hurl

import (
	"os"
	"path/filepath"
	"strings"
)

// CollectFiles returns all .hurl file paths in the given directory.
func CollectFiles(dir string) []string {
	var files []string
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil
	}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".hurl") {
			files = append(files, filepath.Join(dir, e.Name()))
		}
	}
	return files
}

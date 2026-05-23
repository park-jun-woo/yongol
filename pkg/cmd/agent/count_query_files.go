//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what countQueryFiles — db/queries 디렉토리의 파일 수 반환

package agent

import (
	"os"
	"path/filepath"
)

func countQueryFiles(specsDir string) int {
	queriesDir := filepath.Join(specsDir, "db", "queries")
	entries, err := os.ReadDir(queriesDir)
	if err != nil {
		return 0
	}
	n := 0
	for _, e := range entries {
		if !e.IsDir() {
			n++
		}
	}
	return n
}

//ff:func feature=gen-nestjs type=test-helper control=sequence
//ff:what countTSFiles — 디렉토리 트리에서 .ts 파일 개수 카운트 헬퍼
package nestjs

import (
	"os"
	"path/filepath"
)

// countTSFiles walks dir and returns the number of regular .ts files.
func countTSFiles(dir string) int {
	tsFiles := 0
	_ = filepath.Walk(dir, func(p string, info os.FileInfo, e error) error {
		if e == nil && info != nil && !info.IsDir() && filepath.Ext(p) == ".ts" {
			tsFiles++
		}
		return nil
	})
	return tsFiles
}

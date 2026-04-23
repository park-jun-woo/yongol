//ff:func feature=migration type=util control=iteration dimension=1
//ff:what listSQLFiles — 디렉토리에서 *.sql 파일 경로 수집 (skip 목록 제외, 정렬)
package migration

import (
	"fmt"
	"os"
	"path/filepath"
)

// listSQLFiles returns the sorted absolute paths of *.sql files in `dir`
// excluding anything in `skipFiles`. When the directory does not exist
// it returns (nil, nil) so callers can treat it as empty.
func listSQLFiles(dir string, skipFiles []string) ([]string, error) {
	skip := map[string]bool{}
	for _, f := range skipFiles {
		skip[f] = true
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("read dir %s: %w", dir, err)
	}
	var files []string
	for _, e := range entries {
		if name := e.Name(); shouldParseSQL(e.IsDir(), name, skip) {
			files = append(files, filepath.Join(dir, name))
		}
	}
	sortStringSlice(files)
	return files, nil
}

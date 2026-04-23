//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what cleanStaleAuthFiles — authKeepFiles 에 없는 *.go 파일을 제거한다

package auth

import (
	"os"
	"path/filepath"
)

// cleanStaleAuthFiles removes *.go files in authDir that are not in
// authKeepFiles. Directories and non-.go files are ignored.
func cleanStaleAuthFiles(authDir string) error {
	entries, err := os.ReadDir(authDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if filepath.Ext(name) != ".go" {
			continue
		}
		if authKeepFiles[name] {
			continue
		}
		if err := os.Remove(filepath.Join(authDir, name)); err != nil {
			return err
		}
	}
	return nil
}

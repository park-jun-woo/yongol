//ff:func feature=generate type=helper control=sequence
//ff:what dirExists — 디렉토리 존재 여부 확인
package generate

import "os"

func dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

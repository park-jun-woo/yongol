//ff:func feature=cli type=test control=sequence
//ff:what TestStatusCmd — statusCmd 서브커맨드 기본 동작 검증
package main

import (
	"os"
	"path/filepath"
)

func writeTestFile(dir, name, content string) error {
	return os.WriteFile(filepath.Join(dir, name), []byte(content), 0o644)
}

//ff:func feature=gen-gogin type=generator control=sequence
//ff:what writeValidator — request_validator.go 소스를 artifacts 디렉토리에 기록

package middleware

import (
	"os"
	"path/filepath"
)

// writeValidator emits the request_validator.go source.
func writeValidator(mwDir string) error {
	return os.WriteFile(filepath.Join(mwDir, "request_validator.go"), []byte(requestValidatorSource), 0o644)
}

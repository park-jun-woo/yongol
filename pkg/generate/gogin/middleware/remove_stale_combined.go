//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what removeStaleCombined -- pre-split 합본 파일(prometheus.go 등) 잔존 시 삭제

package middleware

import (
	"os"
	"path/filepath"
)

// staleCombinedFiles lists filenames from previous codegen runs that have
// since been replaced by per-func split files.  Removing them prevents
// filefunc F1 violations caused by leftover multi-func files.
var staleCombinedFiles = []string{
	"prometheus.go",       // → prometheus_middleware.go + prometheus_handler.go
	"rate_limit.go",       // → fixed_rate_limit.go + fixed_rate_limit_key.go
	"security_headers.go", // → 6 split files
}

// removeStaleCombined deletes known pre-split combined files from mwDir.
// Missing files are silently ignored.
func removeStaleCombined(mwDir string) error {
	for _, name := range staleCombinedFiles {
		p := filepath.Join(mwDir, name)
		if err := os.Remove(p); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

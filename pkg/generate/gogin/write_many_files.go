//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what WriteManyFiles — 같은 디렉토리에 여러 파일을 쓰고 생성된 파일명을 반환

package gogin

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ffhash"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
)

// WriteManyFiles writes each (name → content) pair to dir with 0644
// permissions, creating dir when absent. It returns the set of file
// names that were written so callers can hand the set to CleanStaleFiles
// for post-write sweeping.
//
// Writes stop at the first error; partially written files remain on disk
// by design so a re-run of the generator can observe the failure instead
// of silently rolling back. The dir argument is created with MkdirAll so
// nested paths are supported.
func WriteManyFiles(dir string, files map[string]string) (map[string]bool, error) {
	written := make(map[string]bool, len(files))
	if len(files) == 0 {
		return written, nil
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return written, err
	}
	for name, body := range files {
		path := filepath.Join(dir, name)
		content := []byte(body)
		if strings.HasSuffix(name, ".go") {
			content = ffhash.InjectCheckedLine(content)
		}
		if err := fffile.WriteIfNotPreserved(path, content); err != nil {
			return written, err
		}
		written[name] = true
	}
	return written, nil
}

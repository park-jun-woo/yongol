//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what writeFiles — 같은 디렉토리에 여러 파일을 쓰고 //ff:checked 주입

package middleware

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ffhash"
	"github.com/park-jun-woo/yongol/pkg/generate/gogin/fffile"
)

// writeFiles writes each (name, content) pair to dir, injecting
// //ff:checked lines into .go files. Uses fffile.WriteIfNotPreserved so
// user-checked files are preserved.
func writeFiles(dir string, files map[string]string) error {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}
	for name, body := range files {
		path := filepath.Join(dir, name)
		content := []byte(body)
		if strings.HasSuffix(name, ".go") {
			content = ffhash.InjectCheckedLine(content)
		}
		if err := fffile.WriteIfNotPreserved(path, content); err != nil {
			return err
		}
	}
	return nil
}

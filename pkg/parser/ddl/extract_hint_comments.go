//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what ExtractHintCommentsFromDir — 디렉토리 내 *.sql 파일을 순회하며 yongol 힌트 주석 수집

package ddl

import (
	"os"
	"path/filepath"
	"strings"
)

// ExtractHintCommentsFromDir walks <dir> for *.sql files and returns
// every yongol hint comment it finds. Unknown `-- @foo` tags are
// ignored so existing projects keep working.
func ExtractHintCommentsFromDir(dir string) ([]HintComment, error) {
	var out []HintComment
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(strings.ToLower(e.Name()), ".sql") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		hints, err := scanHintComments(f, path)
		f.Close()
		if err != nil {
			return nil, err
		}
		out = append(out, hints...)
	}
	return out, nil
}

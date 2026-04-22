//ff:func feature=tsx-parser type=command control=iteration dimension=1
//ff:what frontend 디렉토리 아래 .tsx 파일 전체를 파싱
package tsx

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strings"
)

// ParseDir walks root recursively and parses every .tsx file found. Returns
// the aggregate PageSpec list and a best-effort aggregated error (one error
// per failing file, joined with '; ').
//
// Files under node_modules/, dist/, build/ and hidden dirs are skipped so
// generated / vendored TSX never pollutes the SSOT set.
func ParseDir(root string) ([]PageSpec, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			name := d.Name()
			if name == "node_modules" || name == "dist" || name == "build" || (strings.HasPrefix(name, ".") && path != root) {
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, ".tsx") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Strings(files)
	var pages []PageSpec
	var errs []string
	for _, f := range files {
		p, err := Parse(f)
		if err != nil {
			errs = append(errs, err.Error())
			continue
		}
		pages = append(pages, p)
	}
	if len(errs) > 0 {
		return pages, fmt.Errorf("%d tsx parse error(s): %s", len(errs), strings.Join(errs, "; "))
	}
	return pages, nil
}

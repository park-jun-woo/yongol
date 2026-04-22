//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 디렉토리 내 모든 .ssac 파일을 재귀 탐색하여 []ServiceFunc 반환
package ssac

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// ParseDir은 디렉토리 내 모든 .ssac 파일을 재귀 탐색하여 []ServiceFunc를 반환한다.
func ParseDir(dir string) ([]ServiceFunc, []diagnostic.Diagnostic) {
	var funcs []ServiceFunc
	var diags []diagnostic.Diagnostic
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			diags = append(diags, diagnostic.Diagnostic{
				File:    path,
				Line:    0,
				Phase:   diagnostic.PhaseParse,
				Level:   diagnostic.LevelError,
				Message: "디렉토리 탐색 실패: " + err.Error(),
			})
			return nil
		}
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".ssac") {
			return nil
		}
		parsed, d2 := parseDirEntry(dir, path, d.Name())
		diags = append(diags, d2...)
		funcs = append(funcs, parsed...)
		return nil
	})
	return funcs, diags
}

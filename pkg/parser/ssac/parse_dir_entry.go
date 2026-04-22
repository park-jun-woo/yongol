//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what 단일 디렉토리 엔트리의 .ssac 파일을 파싱하여 feature 할당
package ssac

import (
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// parseDirEntry는 단일 .ssac 파일을 파싱하고 feature를 할당한다.
func parseDirEntry(dir, path, name string) ([]ServiceFunc, []diagnostic.Diagnostic) {
	sfs, diags := ParseFile(path)
	if len(diags) > 0 {
		return nil, diags
	}
	rel, _ := filepath.Rel(dir, path)
	if filepath.Dir(rel) == "." {
		return nil, []diagnostic.Diagnostic{{
			File:    path,
			Line:    0,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: name + " — service/ 직접에 SSaC 파일을 둘 수 없습니다. feature 서브 폴더를 사용하세요 (예: service/auth/" + name + ")",
		}}
	}
	for i := range sfs {
		parts := strings.Split(filepath.Dir(rel), string(filepath.Separator))
		sfs[i].Feature = parts[0]
	}
	return sfs, nil
}

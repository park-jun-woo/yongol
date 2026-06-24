//ff:func feature=generate type=util control=sequence
//ff:what frontendDirExists — 단일/도메인 모드 모두에서 프론트엔드 산출물 디렉토리 존재 여부 확인
package generate

import (
	"path/filepath"
	"slices"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func frontendDirExists(fs *yongol.Fullstack, artifactsDir string) bool {
	if fs != nil && fs.IsDomained() {
		return slices.ContainsFunc(fs.DomainNames(), func(name string) bool {
			return dirExists(filepath.Join(artifactsDir, "frontend", name))
		})
	}
	return dirExists(filepath.Join(artifactsDir, "frontend"))
}

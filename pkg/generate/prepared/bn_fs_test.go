//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func bnFS(mc *pmanifest.ProjectConfig, funcs []ssac.ServiceFunc) *yongol.Fullstack {
	return &yongol.Fullstack{Manifest: mc, ServiceFuncs: funcs}
}

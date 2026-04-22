//ff:func feature=gen-filefunc type=util control=iteration dimension=1
//ff:what addFuncSpecFeatures — Func 스펙의 Package 이름을 Description 과 함께 맵에 추가
package filefunc

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// addFuncSpecFeatures inserts funcspec package names into dst using the
// spec's Description as initial value. Existing entries with a non-empty
// description are preserved.
func addFuncSpecFeatures(dst map[string]string, fs *yongol.Fullstack) {
	for i := range fs.ProjectFuncSpecs {
		pkg := strings.TrimSpace(fs.ProjectFuncSpecs[i].Package)
		desc := strings.TrimSpace(fs.ProjectFuncSpecs[i].Description)
		upsertFeatureDesc(dst, pkg, desc)
	}
}

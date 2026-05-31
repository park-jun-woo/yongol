//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeCrossFeatureModuleRefs — imports 배열에 cross-feature Module 참조 항목 출력

package ssac

import (
	"fmt"
	"strings"
)

// writeCrossFeatureModuleRefs writes one module reference entry per cross-feature
// dependency inside the @Module imports array.
func writeCrossFeatureModuleRefs(b *strings.Builder, crossFeatures []string) {
	for _, cf := range crossFeatures {
		modName := strings.ToUpper(cf[:1]) + cf[1:] + "Module"
		b.WriteString(fmt.Sprintf("    %s,\n", modName))
	}
}

//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeCrossFeatureImports — cross-feature module import 문 출력

package ssac

import (
	"fmt"
	"strings"
)

// writeCrossFeatureImports writes one module import statement per cross-feature
// dependency.
func writeCrossFeatureImports(b *strings.Builder, crossFeatures []string) {
	for _, cf := range crossFeatures {
		modName := strings.ToUpper(cf[:1]) + cf[1:] + "Module"
		b.WriteString(fmt.Sprintf("import { %s } from '../%s/%s.module';\n", modName, cf, cf))
	}
}

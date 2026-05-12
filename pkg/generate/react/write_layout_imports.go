//ff:func feature=gen-react type=emitter control=iteration dimension=1
//ff:what 레이아웃 컴포넌트 import 문을 Builder에 기록한다

package react

import (
	"fmt"
	"strings"
)

// writeLayoutImports writes layout component import statements.
func writeLayoutImports(sb *strings.Builder, layoutNames []string) {
	for _, name := range layoutNames {
		if name == "" {
			continue
		}
		compName := layoutComponentName(name)
		fmt.Fprintf(sb, "import %s from './layouts/%s'\n", compName, compName)
	}
}

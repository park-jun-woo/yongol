//ff:func feature=gen-react type=emitter control=iteration dimension=1
//ff:what DESIGN.md 추가 색상 토큰을 Builder에 기록한다

package react

import (
	"fmt"
	"strings"
)

// writeExtraDesignColors emits DESIGN.md color tokens not already covered
// by the semantic slots.
func writeExtraDesignColors(b *strings.Builder, colors map[string]string) {
	keys := sortedMapKeys(colors)
	for _, k := range keys {
		if semanticColorNames[k] {
			continue
		}
		b.WriteString(fmt.Sprintf("        '%s': '%s',\n", k, colors[k]))
	}
}

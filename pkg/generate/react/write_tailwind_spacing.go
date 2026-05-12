//ff:func feature=gen-react type=emitter control=iteration dimension=1
//ff:what tailwind.config.js의 spacing 섹션을 Builder에 기록한다

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// writeTailwindSpacing writes the spacing section of tailwind.config.js.
func writeTailwindSpacing(b *strings.Builder, dspec *design.DesignSpec) {
	if dspec == nil || len(dspec.Spacing) == 0 {
		return
	}
	b.WriteString("      spacing: {\n")
	keys := sortedMapKeys(dspec.Spacing)
	for _, k := range keys {
		b.WriteString(fmt.Sprintf("        %s: '%s',\n", k, dspec.Spacing[k]))
	}
	b.WriteString("      },\n")
}

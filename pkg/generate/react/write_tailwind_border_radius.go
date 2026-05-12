//ff:func feature=gen-react type=emitter control=sequence
//ff:what tailwind.config.js의 borderRadius 섹션을 Builder에 기록한다

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// writeTailwindBorderRadius writes the borderRadius section of tailwind.config.js.
func writeTailwindBorderRadius(b *strings.Builder, theme *manifest.FrontendTheme, dspec *design.DesignSpec) {
	b.WriteString("      borderRadius: {\n")
	if dspec != nil && len(dspec.Rounded) > 0 {
		keys := sortedMapKeys(dspec.Rounded)
		for _, k := range keys {
			b.WriteString(fmt.Sprintf("        %s: '%s',\n", k, dspec.Rounded[k]))
		}
	} else {
		b.WriteString(fmt.Sprintf("        lg: '%s',\n", orDefault(theme, pickRadius, "0.5rem")))
		b.WriteString("        md: 'calc(var(--radius, 0.5rem) - 2px)',\n")
		b.WriteString("        sm: 'calc(var(--radius, 0.5rem) - 4px)',\n")
	}
	b.WriteString("      },\n")
}

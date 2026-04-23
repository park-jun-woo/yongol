//ff:func feature=gen-react type=util control=sequence
//ff:what writeColorToken — tailwind colors.<name> 에 { DEFAULT, foreground } 쌍 기록

package react

import (
	"fmt"
	"strings"
)

// writeColorToken emits a {DEFAULT, foreground} pair under `colors.<name>`.
// shadcn conventions expect both so `bg-<name>` and `text-<name>-foreground`
// work out of the box.
func writeColorToken(b *strings.Builder, name, def, foreground string) {
	b.WriteString(fmt.Sprintf("        %s: { DEFAULT: '%s', foreground: '%s' },\n", name, def, foreground))
}

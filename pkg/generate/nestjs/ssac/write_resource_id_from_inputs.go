//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeResourceIDFromInputs — Inputs 중 ResourceID 항목을 resourceId 인자로 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeResourceIDFromInputs writes the resourceId entry derived from a
// ResourceID entry in inputs, when present.
func writeResourceIDFromInputs(b *strings.Builder, inputs []ir.FieldArg, indent string) {
	for _, input := range inputs {
		if input.Key == "ResourceID" {
			b.WriteString(fmt.Sprintf("%s  resourceId: String(%s),\n", indent, renderArgValue(input)))
		}
	}
}

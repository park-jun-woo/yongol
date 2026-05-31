//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what writeAuthInputs — authz.check 호출 인자(ResourceID 제외) key: value 줄 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeAuthInputs writes the authz.check object entries, skipping the
// ResourceID input which is handled separately.
func writeAuthInputs(b *strings.Builder, inputs []ir.FieldArg, indent string) {
	for _, input := range inputs {
		if input.Key == "ResourceID" {
			continue
		}
		b.WriteString(fmt.Sprintf("%s  %s: %s,\n", indent, toSnake(input.Key), renderArgValue(input)))
	}
}

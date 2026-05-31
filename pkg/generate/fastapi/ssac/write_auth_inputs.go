//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what writeAuthInputs — authz_check 호출 인자(ResourceID 제외) 키=값 줄 출력

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeAuthInputs writes the authz_check keyword arguments, skipping the
// ResourceID input which is handled separately.
func writeAuthInputs(b *strings.Builder, inputs []ir.FieldArg, indent string) {
	for _, input := range inputs {
		if input.Key == "ResourceID" {
			continue
		}
		b.WriteString(fmt.Sprintf("%s    %s=%s,\n", indent, resolveArgKey(input), renderArgValue(input)))
	}
}

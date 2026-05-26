//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderPublishOp — PublishOp → queue.publish TypeScript await 문 렌더링

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderPublishOp writes a queue publish call.
func renderPublishOp(b *strings.Builder, op *ir.PublishOp, indent string) {
	if op == nil {
		return
	}
	b.WriteString(fmt.Sprintf("%sawait this.queue.publish('%s', {\n", indent, op.Topic))
	for _, p := range op.Payload {
		b.WriteString(fmt.Sprintf("%s  %s: %s,\n", indent, p.Key, renderArgValue(p)))
	}
	b.WriteString(fmt.Sprintf("%s});\n", indent))
}
